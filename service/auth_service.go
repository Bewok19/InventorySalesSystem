package service

import (
	"errors"
	"log"
	"myapp/entity"
	"myapp/repository"
	"myapp/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
    Register(request entity.RegisterRequest) (*entity.User, error)
    Login(email, password string) (string, error)
}


type authServiceImpl struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authServiceImpl{userRepository: userRepo}
}

func (s *authServiceImpl) VerifyPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func (s *authServiceImpl) GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := "your_secret_key"
	return token.SignedString([]byte(secretKey))
}

func (s *authServiceImpl) Register(request entity.RegisterRequest) (*entity.User, error) {
    // Periksa apakah email sudah terdaftar
    existingUser, err := s.userRepository.FindByEmail(request.Email)
    if err != nil {
        return nil, err
    }
    if existingUser != nil {
        return nil, errors.New("email already registered")
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(request.Password)
    if err != nil {
        return nil, err
    }

    // Buat user baru
    newUser := &entity.User{
        Username:     request.Name,
        Email:    request.Email,
        Password: hashedPassword,
        Role:     request.Role,
    }

    // Simpan user ke database
    err = s.userRepository.Save(newUser)
    if err != nil {
        return nil, err
    }

    return newUser, nil
}


// Login handles user authentication
func (s *authServiceImpl) Login(email, password string) (string, error) {
	// Validate inputs
	if email == "" || password == "" {
		log.Println("[ERROR] Email or password is empty")
		return "", errors.New("email and password are required")
	}

	// Find user by email
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		log.Printf("[ERROR] User not found: %s, error: %v", email, err)
		return "", errors.New("invalid credentials")
	}
	if user == nil {
		log.Printf("[WARN] User not found: %s", email)
		return "", errors.New("invalid credentials")
	}

	// Verify password
	if !utils.VerifyPassword(user.Password, password) {
        log.Println("Invalid password for user:", email)
        return "", errors.New("invalid credentials")
    }
    

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Role)
if err != nil {
    log.Println("Error generating token:", err)
    return "", err
}


	log.Printf("[INFO] Login successful for user: %s", email)
	return token, nil
}
