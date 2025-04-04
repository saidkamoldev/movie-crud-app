package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-crud-app/internal/repository/models"
	"strconv"
)

type MovieHandler struct {
	DB *gorm.DB
}

func NewMovieHandler(db *gorm.DB) *MovieHandler {
	return &MovieHandler{DB: db}
}

// GetMovies godoc
// @Summary Get all movies
// @Description Get a list of all movies from the database
// @Tags Movies
// @Accept json
// @Produce json
// @Success 200 {array} models.Movie "List of movies"
// @Failure 500 {object} gin.H{"error": "Database error"}
// @Failure 404 {object} gin.H{"message": "No movies found"}
// @Router /movies [get]
func (h *MovieHandler) GetMovies(c *gin.Context) {
	var movies []models.Movie
	if err := h.DB.Find(&movies).Error; err != nil {
		c.JSON(500, gin.H{"error": "Database error: " + err.Error()})
		return
	}
	if len(movies) == 0 {
		c.JSON(404, gin.H{"message": "No movies found"})
		return
	}
	c.JSON(200, movies)
}

// CreateMovie godoc
// @Summary Create a new movie
// @Description Create a new movie and save it to the database
// @Tags Movies
// @Accept json
// @Produce json
// @Param movie body models.Movie true "Movie data"
// @Success 201 {object} models.Movie "Created movie"
// @Failure 400 {object} gin.H{"error": "Invalid data"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /movies [post]
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	if movie.Year <= 0 {
		c.JSON(400, gin.H{"error": "Year must be a valid positive integer"})
		return
	}
	if err := h.DB.Create(&movie).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, movie)
}

// UpdateMovie godoc
// @Summary Update an existing movie
// @Description Update a movie in the database
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param movie body models.Movie true "Updated movie data"
// @Success 200 {object} models.Movie "Updated movie"
// @Failure 400 {object} gin.H{"error": "Invalid data"}
// @Failure 404 {object} gin.H{"error": "Movie not found"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /movies/{id} [put]
func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie

	if err := h.DB.First(&movie, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Movie not found"})
		return
	}

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data: " + err.Error()})
		return
	}

	if movie.Title == "" || movie.Director == "" || movie.Year == 0 {
		c.JSON(400, gin.H{"error": "Title, Director, and Year are required"})
		return
	}

	if _, err := strconv.Atoi(strconv.Itoa(movie.Year)); err != nil {
		c.JSON(400, gin.H{"error": "Year must be a valid integer"})
		return
	}

	if err := h.DB.Save(&movie).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, movie)
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a movie by its ID from the database
// @Tags Movies
// @Param id path int true "Movie ID"
// @Success 200 {object} gin.H{"message": "Movie deleted"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /movies/{id} [delete]
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	if err := h.DB.Delete(&models.Movie{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Movie deleted"})
}
