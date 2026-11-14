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
	Image       string
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
		"Title": "Hari Nandan — Software Developer",
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
				Title: "Mathrubhumi.com — Qikz", Period: "MAY 2024 — AUG 2024", Type: "INTERNSHIP · WEB APPLICATION",
				Description: "Worked on Qikz web app features and real-time engagement tools for a news platform.",
				Impact: "Contributed production features and built a real-time election result widget used on the homepage.",
				Stack: "Next.js, JavaScript, Web APIs", Role: "Intern developer",
				Image: "/static/projects/mathrubhumi-qikz.png",
			},
			{
				Title: "Namaste Service — Authorized Service Center Platform", Period: "PENTIUMLABS-PROJECT", Type: "SERVICE MANAGEMENT · WEB PLATFORM",
				Description: "A platform for authorized service centers, designed to support their service operations.",
				Impact: "Brings authorized service center workflows onto a web platform.",
				Stack: "Web platform", Role: "Software development",
				Image: "/static/projects/namaste-service.png",
			},
			{
				Title: "ServiceBill — SaaS Billing Platform", Period: "PENTIUMLABS-PROJECT", Type: "SAAS · BACKEND ENGINEERING",
				Description: "Designed and developed a multi-tenant SaaS billing system with pay-as-you-go pricing, an admin portal and mobile integration.",
				Impact: "Used by authorized service centres to manage billing operations and customer subscriptions.",
				Stack: "Go, Modular Monolith, PostgreSQL, Docker, Dokploy", Role: "Backend + deployment engineer · Team of 3",
				Image: "/static/projects/servicebill.png",
			},
			{
				Title: "SalesGo — Delivery Management System", Period: "FREELANCE", Type: "BACKEND · INFRASTRUCTURE",
				Description: "Built the backend for delivery tracking, driver allocation and customer management for a trading company.",
				Impact: "In production at Growmore Trading, supporting around 40 drivers and 3,000+ customers.",
				Stack: "Node.js, Microservices, MongoDB, AWS, Docker, Nginx", Role: "Backend + deployment engineer · Team of 3",
				Image: "/static/projects/salesgo.png",
			},
			{
				Title: "RKM Class Grading App", Period: "FREELANCE", Type: "MOBILE APPLICATION",
				Description: "Developed a mobile app for RKM School, Kozhikode, helping teachers manage and evaluate class marks.",
				Impact: "Actively used by 50+ teachers for grading and internal assessments.",
				Stack: "Capacitor.js, Firebase", Role: "Solo developer",
				Image: "/static/projects/rkm-class-grading.png",
			},
			{
				Title: "OpenDentalCAD", Period: "OPEN PROJECT", Type: "DENTAL · WEB APPLICATION",
				Description: "A web-based dental software application.",
				Impact: "Provides dental software workflows through a web-based application.",
				Stack: "Web application", Role: "Software development",
				Image: "/static/projects/opendentalcad.png",
			},
			{
				Title: "Custom AI Resume Builder", Period: "OPEN PROJECT", Type: "AI · WEB APPLICATION",
				Description: "An open-source custom resume builder that uses AI to tailor a resume to a specific job, align skills with job requirements and improve ATS compatibility. https://customresume.duckdns.org",
				Impact: "Freely used by multiple users.",
				Stack: "JavaScript, Go, Docker", Role: "Full-stack development",
				Image: "/static/projects/custom-ai-resume-builder.png",
			},
		},
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template Execution Error: "+err.Error(), http.StatusInternalServerError)
	}
}
