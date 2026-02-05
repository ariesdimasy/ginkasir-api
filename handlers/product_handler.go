package handlers

import (
	"ginkasir/models"
	"ginkasir/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) SetupRoutes(router *gin.RouterGroup) {
	productGroup := router.Group("/products")
	{
		productGroup.GET("/", h.GetAll)
		productGroup.GET("/:id", h.GetbyID)
		productGroup.POST("/", h.Create)
		productGroup.PUT("/:id", h.Update)
		productGroup.DELETE("/:id", h.Delete)
	}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	name := ctx.Query("name")

	products, total, errPs := h.service.GetAllProducts(&models.SearchProductRequest{
		Page:  page,
		Limit: limit,
		Name:  name,
	})

	if errPs != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": errPs.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "get products success",
		"data":    products,
		"meta": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})

}

func (h *ProductHandler) GetbyID(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	product, err := h.service.GetProductByID(int64(id))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "get product success",
		"data":    product,
	})

}

func (h *ProductHandler) Create(ctx *gin.Context) {

	var req *models.CreateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	product := h.service.CreateProduct(req)

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "create product success",
		"data":    product,
	})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	var req *models.UpdateProductRequest
	var id, _ = strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	product := h.service.UpdateProduct(int64(id), req)

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "update product success",
		"data":    product,
	})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {

	var id, _ = strconv.Atoi(ctx.Param("id"))

	product := h.service.DeleteProduct(int64(id))

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "delete product success",
		"data":    product,
	})
}
