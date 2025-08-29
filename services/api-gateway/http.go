package main

import (
	"gcjade/services/api-gateway/grpc_clients"
	"gcjade/shared/contracts"
	pb "gcjade/shared/proto/catalogue"
	"log"
	"net/http"
)

func handleCreateCategory(w http.ResponseWriter, r *http.Request) {
	var payload createCategoryPayload

	if err := readJSON(w, r, &payload); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	catalogueService, err := grpc_clients.NewCatalogueServiceClient()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create catalog service client", http.StatusInternalServerError)
		return
	}
	defer catalogueService.Close()

	res, err := catalogueService.Client.CreateCategory(r.Context(), &pb.CreateCategoryRequest{
		Name:        payload.Name,
		Description: payload.Description,
	})

	if err != nil {
		http.Error(w, "failed to create category", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, contracts.APIResponse{
		Data: res,
	})

}

func handleListCategories(w http.ResponseWriter, r *http.Request) {
	catalogueService, err := grpc_clients.NewCatalogueServiceClient()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create catalog service client", http.StatusInternalServerError)
		return
	}
	defer catalogueService.Close()

	res, err := catalogueService.Client.ListCategories(r.Context(), &pb.ListCategoriesRequest{})
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to list categories", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, contracts.APIResponse{
		Data: res.Categories,
	})
}

func handleFindCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	log.Println("ID: ", id)

	catalogueService, err := grpc_clients.NewCatalogueServiceClient()
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create catalog service client", http.StatusInternalServerError)
		return
	}
	defer catalogueService.Close()

	res, err := catalogueService.Client.FindCategoryByID(r.Context(), &pb.FindCategoryByIDRequest{
		Id: id,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJSON(w, http.StatusCreated, contracts.APIResponse{
		Data: res,
	})

}
