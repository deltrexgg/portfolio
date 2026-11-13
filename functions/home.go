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
				Title: "Custom AI Resume Builder", Period: "PRESENT", Type: "AI · WEB APPLICATION",
				Description: "An open-source custom resume builder that uses AI to tailor a resume to a specific job, align skills with job requirements and improve ATS compatibility. https://customresume.duckdns.org",
				Impact: "Freely used by multiple users.",
				Stack: "JavaScript, Go, Docker", Role: "Full-stack development",
				Image: "/static/projects/custom-ai-resume-builder.png",
			},
			{
				Title: "ServiceBill — SaaS Billing Platform", Period: "FREELANCE · PRESENT", Type: "SAAS · BACKEND ENGINEERING",
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
				Title: "Mathrubhumi.com — Qikz", Period: "MAY 2024 — AUG 2024", Type: "INTERNSHIP · WEB APPLICATION",
				Description: "Worked on Qikz web app features and real-time engagement tools for a news platform.",
				Impact: "Contributed production features and built a real-time election result widget used on the homepage.",
				Stack: "Next.js, JavaScript, Web APIs", Role: "Intern developer",
				Image: "/static/projects/mathrubhumi-qikz.png",
			},
		},
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template Execution Error: "+err.Error(), http.StatusInternalServerError)
	}
}
