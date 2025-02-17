package main

import (
	"fmt"

	"github.com/USACE/go-consequences/compute"
	"github.com/USACE/go-consequences/hazardproviders"
	"github.com/USACE/go-consequences/resultswriters"
	"github.com/USACE/go-consequences/structureprovider"
)

func main() {
	fmt.Println("Start")

	//define flood hazard, provide the file path to your desired flood
	hazard, e := hazardproviders.Init("./dev/go-consequences/data/flood_here.tif")
	if e != nil {
		panic(e)
	}
	defer hazard.Close()

	//define structure inventory, provide file path to structure inventory
	inventory, b := structureprovider.InitSHP("./dev/go-consequences/data/inventory_here.shp")
	if b != nil {
		panic(b)
	}

	//define result file, provide file path where you want results to be saved
	// results, c := resultswriters.InitGpkResultsWriter("./output/result_file.gpkg", "event_results")
	results, c := resultswriters.InitGpkResultsWriter_Projected("./dev/go-consequences/output/result_file.gpkg", "event_results", 9822)
	if c != nil {
		panic(c)
	}
	defer results.Close()

	//run the compute
	compute.StreamAbstract(hazard, inventory, results)
	fmt.Println("End")
}