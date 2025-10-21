// internal/virtual/demo_merchants.go
package virtual

type DemoMerchant struct {
    ID          string
    Name        string
    Country     string
    Category    string
    Description string
    Products    []DemoProduct
}

type DemoProduct struct {
    Name        string
    Description string
    Price       int64 // JUC
}

var DemoMerchants = []DemoMerchant{
    {
        ID:          "demo_tech_ke",
        Name:        "Nairobi Tech Hub",
        Country:     "KE",
        Category:    "Technology",
        Description: "Virtual electronics and gadgets store",
        Products: []DemoProduct{
            {Name: "Virtual Smartphone", Description: "Latest model smartphone", Price: 800},
            {Name: "Virtual Laptop", Description: "High-performance laptop", Price: 1500},
            {Name: "Virtual Headphones", Description: "Noise-cancelling headphones", Price: 200},
        },
    },
    {
        ID:          "demo_fashion_ng", 
        Name:        "Lagos Fashion Store",
        Country:     "NG",
        Category:    "Fashion",
        Description: "Virtual clothing and accessories",
        Products: []DemoProduct{
            {Name: "Virtual Designer Dress", Description: "Elegant evening dress", Price: 300},
            {Name: "Virtual Sneakers", Description: "Limited edition sneakers", Price: 250},
            {Name: "Virtual Handbag", Description: "Designer leather handbag", Price: 400},
        },
    },
    {
        ID:          "demo_services_gh",
        Name:        "Accra Digital Services",
        Country:     "GH", 
        Category:    "Services",
        Description: "Virtual consulting and digital services",
        Products: []DemoProduct{
            {Name: "Virtual Website Design", Description: "Custom website development", Price: 1200},
            {Name: "Virtual Marketing Consult", Description: "1-hour marketing strategy", Price: 150},
            {Name: "Virtual SEO Audit", Description: "Comprehensive SEO analysis", Price: 300},
        },
    },
}
