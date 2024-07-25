
package models

import (
	"CVGolang/assets"
	"CVGolang/utils"
	"CVGolang/views"
	"github.com/charmbracelet/lipgloss"
)

func getPageContent() []string {
	return []string{
		getWelcomePage(),
		getExperiencePage(),
		getSkillsPage(),
		getProjectsPage(),
		getHobbiesPage(),
		getSocialsPage(),
		getHelpPage(),
	}
}

func getWelcomePage() string {
	return assets.StefArt1 + "\n\n" +  
  assets.StefArt2 + 
  "\n\nWelcome to my CV (in the terminal!)" + 
  "\n\nI'm Stefan Vuckovic, a 4th Year MEng Student studing Electronics and Software" + 
  "\n\nMy core interests lie in lower-level/embedded programming, as well as machine learning (a fun mix)" + 
  "\n\nSome Technologies I'm Learning about right now:" + 
  "\n - CUDA, GPU Programming" +
  "\n - RNNs & LSTMs in Machine Learning" +
  "\n - Learning Golang, and basic Zig" +
  "\n\nI'm Currently Reading: Programming Massively Parallel Processors (4th Edition)" +
  "\n\nI hope you enjoy this project as much as I enjoyed making it" 
}

func getExperiencePage() string {
	return assets.ExperienceArt + "\n\n" +

		utils.RenderCard("Software Development Intern",
			"Integrated Environmental Solutions Limited",
			"Jun 2024 - Sep 2024",
			"Glasgow, Scotland",
			"As part of the Core Simulation team, I develop and maintain build pipelines, bringing standardisation to all our production repos with extra Integration and Regression testing. \nI also work on refactors on parts of our up to 40 year old codebases",
		) +
    "\n" +
		utils.RenderCard("Freelance Software Engineer",
			"Vectofy (Orange Matter Ltd)",
			"Oct 2023 - Mar 2024",
			"(Remote) Glasgow, Scotland",
			"Along with other students I developed an Excel Add-On for Vectofy that helps companies in the finance industry easily onboard new employees, through a tool that easily, and visually helps them understand how and what cells interact with each other on the Excel Workbook." + 
      "\n\nAs part of this project I led DevOps, creating and maintaining our pipeline, and ensuring continuous deployment of our product. I also worked in a team of 2 for the Back-end Code that linked our React Project to the Excel Sheet using APIs.",
		)  +
    "\n" +
		utils.RenderCard("Electronic Engineer",
			"UGRacing",
			"Oct 2023 - Ongoing",
			"Glasgow, Scotland",
			"As part of the Formula Student team for the University I worked on the FSUK Winning Concept team, as well as doing extra work for the main FS team" +
      "\n\n On the concept car I am developing an Embedded Telemetry System, that encompasses all the information around the car, to be transmitted and visualised on software I made to be used on any computer, anywhere in the world, with on average half a second of latency" + 
      "\n\n For the production car I develop and maintain C/C++ for a in-wheel screen to display car information to the driver in real time" + 
      "\n\n In addition to this I also provide technical assistance to other members of the team, such as working with https://www.embeduk.com/ on the car’s ECU.",
		) 
}

func getSkillsPage() string {
	return assets.SkillsArt + "\n\n" + utils.RenderSkillsAndFrameworks()
}

func getProjectsPage() string {
	return assets.ProjectArt + "\n\n" +
		utils.RenderCard("CV in the Terminal", "A TUI Emulating a Portfolio Website", "Jun 2024", "", "Tired of people emulating the terminal on a website for their portfolio, I've decided to emulate a portfolio website in the terminal. I think this shows my skills better than designing a React app as well \n\n This project involves developing a TUI in Go, as well as maintaining a Linux AWS Lightsail Server") + "\n" +
    utils.RenderCard("CUDA Programming", "CUDA C code for GPU Programming", "Jul 2024", "", "I'm reading the massively informative Programming Massively Parallel Processors, and in addition to working on the exercises in the book, I am trying to apply what I'm learning to other code, such as using the GPU to improve the speed of a CNN model I've made") + "\n" +
		utils.RenderCard("Pytorch Snake", "Deep Q-Net AI Agent Designed to Play Snake", "Oct 2023", "", "I'm using Deep Q-Network with PyTorch, a form of reinforcement learning, and I'm specificially implementing the Bellman Equation")
}

func getHobbiesPage() string {
	return assets.HobbiesArt + "\n" + assets.DevArt + `
------------------------------------------------
Under Development!

Singing:
--------
I have been singing for a large portion of my life, taking it to a high level, being trained at the Royal Conservatoire of Scotland for Classical Singing, where in addition to completing up to Grade 8 with Distinction, I would have completed a diploma if it wasn't for COVID cancellations. I also sang with the National Youth Choir of Great Britain for just short of a decade.

I enjoy singing much less formally now, participating in choirs and acapella groups

Music:
------
Beyond singing, I enjoy exploring different genres of music. From classical to modern pop, each genre offers something unique and enriching. Music helps me unwind and keeps me motivated during tough times.
I also completed up to Grade 7 in Piano with Merit

Gaming:
-------
In my free time I enjoy gaming, particularily in:
  • Competitive tactical FPS (First-Person Shooter) games
  • Strategy games

I appreciate these genres for challenge my strategic thinking, reflexes, as well as being an outlet for a competitve spirit. Games like Counter-Strike, Valorant, and Civilization are some of my favorites.

Sports:
-------
I enjoy participating in various sports, including:
  • Skiing 
  • Cycling
  • Volleyball

I hope this gives you a glimpse into the activities and hobbies that I am passionate about. They all play a significant role in my life, keeping me balanced and happy.

`
}

func getSocialsPage() string {
	return assets.SocialsArt + "\n\n" +
		lipgloss.JoinVertical(lipgloss.Left,
			utils.RenderSocial(assets.GithubArt, "https://github.com/StefVuck"),
			utils.RenderSocial(assets.LinkedinArt, "https://linkedin.com/in/stefan-vu%c4%8dkovi%c4%87-b63952286"),
			utils.RenderSocial(assets.EmailArt, "stefanvuck@gmail.com"),
		)
}

func getHelpPage() string {
	return assets.HelpArt + `

---------

This terminal-based CV is designed with the classic Vim keybinds as well as more "intuitive" ones to navigate through the content.

Keybindings:
-----------
  q, esc  : Quit the application
  h, left : Move to the previous section
  l, right: Move to the next section
  k, up   : Scroll up within a section
  j, down : Scroll down within a section


Naturally since it's running in a terminal, mouse clicks are not supported. This is a true coders keyboard-only experience :)

Enjoy navigating through my CV!
`
}

func InitialModel() views.Model {
	return views.Model{
		Index:  0,
		Pages:  getPageContent(),
		Width:  80, // Default width, will be updated
		Height: 24, // Default height, will be updated
		Offset: 0,
	}
}

