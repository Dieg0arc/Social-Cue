package handlers

import (
    "context"
    "net/http"
    "time"
    
    "github.com/labstack/echo/v4"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
    
    "backend-instagram/config"
    "backend-instagram/models"
)

var (
    mongoClient *mongo.Client
    cfg         *config.Config
)

// InitMongo inicializa la conexión a MongoDB
func InitMongo() {
    cfg = config.GetConfig()
    mongoClient = config.GetMongoClient(cfg)
}

// Login maneja la autenticación de usuarios
func Login(c echo.Context) error {
    // Verificar que MongoDB esté inicializado
    if mongoClient == nil {
        InitMongo()
    }
    
    // Parsear la solicitud
    req := new(models.LoginRequest)
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Formato de solicitud inválido",
        })
    }
    
    // Obtener la colección de usuarios
    collection := config.GetCollection(mongoClient, cfg, cfg.CollectionUsers)
    
    // Buscar usuario por email
    var user models.User
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    err := collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.JSON(http.StatusUnauthorized, map[string]string{
                "error": "Credenciales inválidas",
            })
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Error en el servidor",
        })
    }
    
    // Verificar contraseña
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{
            "error": "Credenciales inválidas",
        })
    }
    
    // En una implementación real, generarías un token JWT aquí
    // Por ahora, devolvemos un token ficticio
    return c.JSON(http.StatusOK, models.LoginResponse{
        Token: "token-ejemplo-123456",
        User: models.User{
            ID:        user.ID,
            Username:  user.Username,
            Email:     user.Email,
            CreatedAt: user.CreatedAt,
            UpdatedAt: user.UpdatedAt,
        },
    })
}

// Register maneja el registro de nuevos usuarios
func Register(c echo.Context) error {
    // Verificar que MongoDB esté inicializado
    if mongoClient == nil {
        InitMongo()
    }
    
    // Parsear la solicitud
    req := new(models.RegisterRequest)
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Formato de solicitud inválido",
        })
    }
    
    // Validar datos (implementación básica)
    if req.Username == "" || req.Email == "" || req.Password == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Todos los campos son requeridos",
        })
    }
    
    // Obtener la colección de usuarios
    collection := config.GetCollection(mongoClient, cfg, cfg.CollectionUsers)
    
    // Verificar si el email ya existe
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    count, err := collection.CountDocuments(ctx, bson.M{"email": req.Email})
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Error en el servidor",
        })
    }
    
    if count > 0 {
        return c.JSON(http.StatusConflict, map[string]string{
            "error": "El email ya está registrado",
        })
    }
    
    // Hashear contraseña
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Error al procesar la contraseña",
        })
    }
    
    // Crear nuevo usuario
    now := time.Now()
    newUser := models.User{
        ID:        primitive.NewObjectID(),
        Username:  req.Username,
        Email:     req.Email,
        Password:  string(hashedPassword),
        CreatedAt: now,
        UpdatedAt: now,
    }
    
    // Insertar en la base de datos
    _, err = collection.InsertOne(ctx, newUser)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Error al registrar usuario",
        })
    }
    
    return c.JSON(http.StatusCreated, models.RegisterResponse{
        Message: "Usuario registrado correctamente",
        UserID:  newUser.ID.Hex(),
    })
}