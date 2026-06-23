package functions

import (
	"html/template"
	"net/http"
)

type Work struct {
	Title       string
	Period      string
	Type        string
	Description string
	Impact      string
	Stack       string
	Role        string
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
        "templates/index.html",
        "components/desk.html",
        "components/skills.html",
        "components/work.html",
        "components/contact.html",
        "components/resume.html",
        "components/portfolio.html",
    )
    if err != nil {
        http.Error(w, "Template Parsing Error: "+err.Error(), http.StatusInternalServerError)
        return
    }

	data := map[string]interface{}{
	"Title":      "Hari Nandan Portfolio",
	"PageHeader": "yooh!",
	"About":      "I am Hari Nandan, I build software's",

	"Skills": []string{
		"Languages: JavaScript, TypeScript, Go",
		"Frameworks: React.js, Next.js, React Native (Expo), Gin",
		"Backend: Node.js, REST APIs, gRPC, GORM",
		"Databases: PostgreSQL, MongoDB, Redis",
		"DevOps & Cloud: AWS, Docker, Kubernetes, Nginx, HAProxy, Dokploy",
		"Frontend & UI: Component Architecture, Responsive Design, Figma",
		"Concepts: OOP, Modular Monolithic, Microservices, System Design Basics, Cloud Deployment",
	},

	"MyWorks": []Work{

	{
		Title: "ServiceBill – SaaS Billing Platform",
		Period: "Freelance | Present",
		Type:   "Golang Backend / SaaS",

		Description: "Designed and developed a multi-tenant SaaS billing system with pay-as-you-go pricing, admin portal, and mobile integration.",

		Impact: "Actively used by multiple authorized service centers for managing billing operations and customer subscriptions.",

		Stack: "Go, Modular Monolithic, PostgreSQL, Docker, Dokploy",

		Role: "Backend Engineer + Deployment Engineer (Team of 3)",
	},

	{
		Title: "SalesGo – Delivery Management System",
		Period: "Freelance",
		Type:   "Backend + Infrastructure",

		Description: "Built a full backend system for delivery tracking, driver allocation, and customer management for a trading company.",

		Impact: "Used in production by Growmore Trading, managing ~40 drivers and 3000+ customers in daily operations.",

		Stack: "Node.js, Microservice, MongoDB, AWS, Docker, Nginx",

		Role: "Backend + Deployment Engineer (Team of 3)",
	},

	{
		Title: "RKM Class Grading App",
		Period: "Freelance",
		Type:   "Mobile App",

		Description: "Developed a mobile app for RKM School Kozhikode. Used by teachers to manage and evaluate class marks efficiently.",

		Impact: "Actively used by 50+ teachers for grading and internal assessments.",

		Stack: "Capacitor.js, Firebase",

		Role: "Solo Developer",
	},

	{
		Title: "Mathrubhumi.com – Internship",
		Period: "May 2024 – Aug 2024",
		Type:   "Frontend + Backend (Next.js)",

		Description: "Worked on Qikz web app features and real-time engagement tools for news platform.",

		Impact: "Contributed to production features and built real-time election result widget used on homepage.",

		Stack: "Next.js, JavaScript, Web APIs",

		Role: "Intern Developer",
	},
},
}

	err = tmpl.Execute(w, data)
    if err != nil {
        println("Template Execution Error:", err.Error()) 
    }
}