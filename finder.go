package main

import (
	"fmt"
	"strings"
)

// Define the phonespecs struct
type phonespecs struct {
	CPU string
	GPU string
	RAM string
}

// Define the brand model type
type brandModel map[string]phonespecs


func getphone_specs(brand, model string, brands map[string]brandModel) string {
  if models, ok := brands[brand]; ok {
    
    if specs, ok := models[model]; ok{
      return fmt.Sprintf("Brand: %s\nModel: %s\nCPU: %s\nGPU:%s\nRAM: %s", brand, model, specs.CPU, specs.GPU, specs.RAM)
    } else {
      return "Phone not found."
    }
  }
  return "Brand not found"
}

func main() {
	// Define the map of Samsung models and their specifications
	var brands map[string]brandModel = map[string]brandModel{
		"SAMSUNG": {
			  "S": { "ARM Cortex-A8 Hummingbird",  "PowerVR SGX540",  "512MB"},
        "S2": { "Samsung Exynos 4210",  "ARM Mali-400 MP4",  "1GB"},
        "S3": { "Samsung Exynos 4412/Qualcomm Snapdragon S4 Plus",  "ARM Mali-400 MP4/ATI Adreno 225",  "1GB"},
        "S4": { "Qualcomm Snapdragon 600",  "ATI Adreno 320",  "2GB"},
        "S5": { "Qualcomm Snapdragon 801",  "ATI Adreno 330",  "2GB"},
        "S6": { "Samsung Exynos 7420",  "ARM Mali-T760 MP8",  "3GB"},
        "S7": { "Samsung Exynos 8890", "ARM Mali-T880 MP12",  "4GB"},
        "S8": { "Samsung Exynos 8895/Qualcomm Snapdragon 835", "ARM Mali-G71 MP20/ATI Adreno 540", "4GB"},
        "S9": { "Samsung Exynos 9810/Qualcomm Snapdragon 845",  "ARM Mali-G72 MP18/ATI Adreno 630",  "4GB"},
        "S10": { "Samsung Exynos 9820/Qualcomm Snapdragon 855",  "ARM Mali-G76 MP12/ATI Adreno 640",  "6/8GB"},
        "S20": { "Samsung Exynos 990/Qualcomm Snapdragon 865",  "ARM Mali-G77 MP11/ATI Adreno 650",  "8GB"},
        "S20 FE":{ "Samsung Exynos 990",  "ARM Mali-G77 MP11",  "6/8GB"},
        "S20 FE 5G": { "Qualcomm Snapdragon 865",  "ATI Adreno 650",  "6/8GB"},
        "S21": { "Samsung Exynos 2100/Qualcomm Snapdragon 888",  "ARM Mali-G78 MP14/ATI Adreno 660",  "8GB"},
        "S22": { "Samsung Exynos 2200/Qualcomm Snapdragon 8 Gen 1",  "AMD Xclipse 920 (RDNA2)/ATI Adreno 730",  "8GB"},
        "S23": { "Qualcomm Snapdragon 8 Gen 2 OC Edition",  "ATI Adreno 740",  "8GB"},
        "S21 FE": { "Samsung Exynos 2100/Qualcomm Snapdragon 888",  "ARM Mali-G78/ATI Adreno 660",  "6/8GB"},
        "S23 FE": { "Samsung Exynos 2200/Qualcomm Snapdragon 8 Gen 1", "AMD Xclipse 920 (RDNA2)/ATI Adreno 730",  "8GB"},
        "S24": {"Samsung Exynos 2400/Snapdragon 8 Gen 3 OC Edition", "ATI Adreno 750", "8/12GB"},
        "FOLD":{ "Qualcomm Snapdragon 855",  "ATI Adreno 640",  "12GB"},
        "Z-FOLD 2":{ "Qualcomm Snapdragon 865",  "ATI Adreno 650",  "12GB"},
        "Z-FOLD 3":{ "Qualcomm Snapdragon 888",  "ATI Adreno 660",  "12GB"},
        "Z-FOLD 4":{ "Qualcomm Snapdragon 8+ Gen 1",  "ATI Adreno 730",  "12GB"},
        "Z-FOLD 5":{ "Qualcomm Snapdragon 8 Gen 2 OC Edition",  "ATI Adreno 740",  "12GB"},
        "Z-FLIP":{ "Qualcomm Snapdragon 855+",  "ATI Adreno 640",  "8GB"},
        "Z-FLIP 5G":{ "Qualcomm Snapdragon 865+",  "ATI Adreno 650", "8GB"},
        "Z-FLIP 3":{ "Qualcomm Snapdragon 888",  "ATI Adreno 660",  "8GB"},
        "Z-FLIP 4":{ "Qualcomm Snapdragon 8+ Gen 1",  "ATI Adreno 730",  "8GB"},
        "Z-FLIP 5":{ "Qualcomm Snapdragon 8 Gen 2 OC Edition",  "ATI Adreno 740", "8GB"},
        "A01":{ "Qualcomm Snapdragon 439",  "ATI Adreno 505",  "2GB"},
        "A01 CORE":{ "Mediatek MT6739",  "PowerVR GE8100",  "1/2GB"},
        "A10":{ "Samsung Exynos 7784",  "ARM Mali-G71",  "2/4GB"},
        "A20":{ "Samsung Exynos 7784",  "ARM Mali-G71",  "3GB"},
        "A30":{ "Samsung Exynos 7904",  "ARM Mali-G71 MP2",  "3/4GB"},
        "A40":{ "Samsung Exynos 7904",  "ARM Mali-G71 MP2",  "4GB"},
        "A50":{ "Samsung Exynos 9610",  "ARM Mali-G72 MP3",  "4/6GB"},
        "A60":{ "Qualcomm Snapdragon 675",  "ATI Adreno 612",  "6GB"},
        "A70":{ "Qualcomm Snapdragon 675",  "ATI Adreno 612",  "6/8GB"},
        "A80":{ "Qualcomm Snapdragon 730",  "ATI Adreno 618",  "8GB"},
        "A90 5G":{ "Qualcomm Snapdragon 855",  "ATI Adreno 640",  "6/8GB"},
        "A02":{ "Mediatek MT6739W",  "PowerVR GE8100", "2/4GB"},
        "A02S":{ "Qualcomm Snapdragon 450",  "ATI Adreno 506",  "1/4GB"},
        "A03":{ "Unisoc T606",  "ARM Mali-G57 MP1",  "3/4GB"},
        "A03S":{ "Mediatek Helio P35",  "PowerVR GE8320",  "2/4GB"}, 
        "A03 CORE":{ "Unisoc SC9863A",  "IMG8322",  "2GB"},
        "A11": { "Qualcomm Snapdragon 450",  "ATI Adreno 506",  "6/4GB"},
        "A12": { "Mediatek Helio P35",  "PowerVR GE8320",  "2/6GB"},
        "A13": { "Samsung Exynos 850",  "ARM Mali-G52",  "3/6GB"},
        "A13 5G": { "Mediatek Dimensity 700",  "ARM Mali-G57 MC2",  "4/6GB"},
        "A14": { "Mediatek Helio G80",  "ARM Mali-G52 MC2",  "4/6GB"},
        "A14 5G": { "Samsung Exynos 1330",  "ARM Mali-G68 MP2",  "4/8GB"},
        "A20S": { "Qualcomm Snapdragon 450",  "ATI Adreno 506",  "2/4GB"},
        "A20E": { "Samsung Exynos 7884",  "ARM Mali-G71 MP2",  "3GB"},
        "A21": { "Mediatek Helio P35",  "PowerVR GE8320",  "3GB"},
        "A21S": { "Samsung Exynos 850",  "ARM Mali-G52",  "2/6GB"},
        "A22": { "Mediatek MT6769V",  "ARM Mali-G52 MC2",  "4/6GB"},
        "A22 5G": { "Mediatek Dimensity 700",  "ARM Mali-G57 MC2",  "4/8GB"},
        "A23": { "Qualcomm Snapdragon 680",  "ATI Adreno 610",  "4/8GB"},
        "A23 5G": { "Qualcomm Snapdragon 695 5G",  "ATI Adreno 619",  "4/8GB"},
        "A24": { "Mediatek Helio G99",  "ARM Mali-G57 MC2",  "4/8GB"},
        "A31": { "Mediatek Helio P65",  "ARM Mali-G52 MC2",  "4/8GB"},
        "A32": { "Mediatek MT6769V",  "ARM Mali-G52 MC2",  "4/8GB"},
        "A32 5G": { "Mediatek Dimensity 720",  "ARM Mali-G57 MC3",  "4/8GB"},
        "A33 5G": { "Samsung Exynos 1280",  "ARM Mali-G68",  "4/8GB"},
        "A34": { "Mediatek Dimensity 1080",  "ARM Mali-G68 MC4",  "6/8GB"},
        "A41": { "Mediatek Helio P65",  "ARM Mali-G52 MC2",  "4/8GB"},
        "A42 5G": { "Qualcomm Snapdragon 750",  "ATI Adreno 619",  "4/8GB"},

		},
	}

  var selectedBrand, selectedModel string
  fmt.Print("Enter a brand: ")
  fmt.Scanln(&selectedBrand)
  selectedBrand = strings.ToUpper(selectedBrand)

  fmt.Print("Enter a model: ")
  fmt.Scanln(&selectedModel)
  selectedModel = strings.ToUpper(selectedModel)

  results := getphone_specs(selectedBrand, selectedModel, brands)
  fmt.Print(results)
  fmt.Println("Press a Button to exit: ")
  fmt.Scan()

	}


