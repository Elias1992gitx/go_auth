package controllers

import (
    "JWT-Authentication-go/database"
    "JWT-Authentication-go/models"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
    "fmt"
    "github.com/golang-jwt/jwt/v5"  // Add this import
    "time"      // Add this import
    "strconv"   // Add this import
)

//for production env set in env file (secure position)
const secretKey = "jK6DX#mP9$vL2@nQ8wR4tY7hC3bE5sA"

//register function
func Register(c *fiber.Ctx) error {
    fmt.Println("Received a registration request")

    //Parses the incoming JSON payload into a map[string]string
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body",
        })
    }

    //Check for Existing User
    var existingUser models.User
    if err := database.DB.Where("email = ?", data["email"]).First(&existingUser).Error; err == nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Email already exists",
        })
    }

    //Hash the Password using bcrypt library
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to hash password",
        })
    }

    //Creates a new models.User instance with the parsed data and hashed password
    user := models.User{
        Name:     data["name"],
        Email:    data["email"],
        Password: hashedPassword,
    }
    if err := database.DB.Create(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create user",
        })
    }

    //Return Success Response
    return c.JSON(fiber.Map{
        "message": "User created successfully",
    })
}


//login function
func Login(c *fiber.Ctx) error {
    fmt.Println("Received a login request")

     // Parse request body
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "failed to parse request body",
        })
    }

    // Check if user exists
    var user models.User
    database.DB.Where("email = ?", data["email"]).First(&user)
    if user.ID == 0 {
        fmt.Println("User not found")
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Invalid credentials",
        })
    }

    // Compare passwords
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
    if err != nil {
        fmt.Println("Invalid Password:", err)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Invalid credentials",
        })
    }

     // Generate JWT token
     claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": strconv.Itoa(int(user.ID)), 
        "exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
    })
    token, err := claims.SignedString([]byte(secretKey))
    if err != nil {
        fmt.Println("Error generating token:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to generate token",
        })
    }

    // Set JWT token in cookie
    cookie := fiber.Cookie{
        Name:     "jwt",
        Value:    token,
        Expires:  time.Now().Add(time.Hour * 24), // Expires in 24 hours
        HTTPOnly: true,
        Secure:   true,
    }
    c.Cookie(&cookie)


     // Authentication successful, return success response
     return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
        "message": "Login successful",
    })

}


// user retrieval function
func User(c *fiber.Ctx) error {
    fmt.Println("Request to get user...")

    // Retrieve JWT token from cookie
    cookie := c.Cookies("jwt")

    // Parse JWT token with claims
    token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })

    // Handle token parsing errors
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Unauthorized",
        })
    }

    // Extract claims from token
    claims, ok := token.Claims.(*jwt.MapClaims)
    if !ok {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to parse claims",
        })
    }

    // Extract user ID from claims
    id, _ := strconv.Atoi((*claims)["sub"].(string))
    user := models.User{ID: uint(id)}

    // Query user from database using ID
    database.DB.Where("id =?", id).First(&user)

    // Return user details as JSON response
    return c.JSON(user)
}


// logout function
func Logout(c *fiber.Ctx) error {
    fmt.Println("Received a logout request")

    // Clear JWT token by setting an empty value and expired time in the cookie
    cookie := fiber.Cookie{
        Name:     "jwt",
        Value:    "",
        Expires:  time.Now().Add(-time.Hour), // Expired 1 hour ago
        HTTPOnly: true,
        Secure:   true,
    }
    c.Cookie(&cookie)

    // Return success response indicating logout was successful
    return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
        "message": "Logout successful",
    })
}
