package service

import (
	"nktrr/pfizer/api_gateway_service/config"
	"nktrr/pfizer/api_gateway_service/internal/products/commands"
	"nktrr/pfizer/api_gateway_service/internal/products/queries"
	kafkaClient "nktrr/pfizer/pkg/kafka"
	"nktrr/pfizer/pkg/logger"
	readerService "nktrr/pfizer/reader_service/proto/product_reader"
)

type ProductService struct {
	Commands *commands.ProductCommands
	Queries  *queries.ProductQueries
}

func NewProductService(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer, rsClient readerService.ReaderServiceClient) *ProductService {

	createProductHandler := commands.NewCreateProductHandler(log, cfg, kafkaProducer)
	updateProductHandler := commands.NewUpdateProductHandler(log, cfg, kafkaProducer)
	deleteProductHandler := commands.NewDeleteProductHandler(log, cfg, kafkaProducer)

	getProductByIdHandler := queries.NewGetProductByIdHandler(log, cfg, rsClient)
	searchProductHandler := queries.NewSearchProductHandler(log, cfg, rsClient)

	productCommands := commands.NewProductCommands(createProductHandler, updateProductHandler, deleteProductHandler)
	productQueries := queries.NewProductQueries(getProductByIdHandler, searchProductHandler)

	return &ProductService{Commands: productCommands, Queries: productQueries}
}
