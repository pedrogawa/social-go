package db

import (
	"context"
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/pedrogawa/social-go/internal/store"
)

var blogTitles = []string{
	"Understanding Goroutines: Concurrency Made Simple",
	"Building RESTful APIs with Go's Gin Framework",
	"Effective Error Handling Patterns in Go",
	"Go Interfaces: The Secret to Flexible Code",
	"Testing in Go: From Unit Tests to Integration Tests",
	"Mastering the Context Package in Go",
	"Memory Management and Garbage Collection in Go",
	"Building Microservices with Go and Docker",
	"Dependency Injection Techniques in Go Applications",
	"Go Generics: A Practical Introduction",
	"Working with JSON in Go: Tips and Tricks",
	"Database Access in Go: From SQL to ORMs",
	"Securing Your Go Web Applications",
	"Performance Optimization Techniques for Go Applications",
	"Building CLI Tools with Go",
	"Go Modules: Managing Dependencies the Right Way",
	"Implementing WebSockets in Go",
	"Reflection in Go: When and How to Use It",
	"Building a Real-time Chat Application with Go",
	"Design Patterns in Go: Practical Examples",
}

var blogContents = []string{
	"Goroutines are one of Go's most powerful features. They're lightweight threads managed by the Go runtime, allowing concurrent execution with minimal resources. Unlike OS threads, you can spawn thousands of goroutines without significant performance impact. This concurrency model makes Go perfect for modern applications that need to handle multiple operations simultaneously.",
	"Error handling in Go follows a different pattern than most languages. Instead of exceptions, Go functions often return an error value which must be explicitly checked. While this requires more explicit code, it results in clearer control flow and prevents errors from being silently ignored. Always check your errors!",
	"The context package is essential for controlling goroutines in Go applications. It provides cancelation signals, deadlines, and can carry request-scoped values across API boundaries. Learning to use context.Context effectively is crucial for building robust, cancellable operations in your Go services.",
	"Go interfaces are implicitly implemented, unlike in languages like Java or C#. This means any type that implements the methods defined by an interface automatically satisfies that interface - no explicit declaration needed. This design encourages composition over inheritance and creates more flexible code.",
	"Testing in Go is built right into the standard library with the testing package. Write tests as functions that begin with 'Test' followed by a name starting with a capital letter. The Go test runner will automatically find and execute these functions, making unit testing straightforward and consistent.",
	"Go's strict formatting standards, enforced by tools like gofmt, eliminate debates about code style. All Go code looks familiar regardless of who wrote it, enhancing readability across projects and teams. This standardization is a key part of Go's philosophy of simplicity and clarity.",
	"Memory management in Go is handled by the garbage collector, which has seen significant improvements in each release. Understanding how the GC works can help you write more efficient Go code. Remember that while the GC handles memory cleanup, it's still important to design your data structures with memory usage in mind.",
	"Dependency injection in Go is often simpler than in other languages. Rather than complex frameworks, Go typically uses simple function parameters or struct fields to inject dependencies. This approach aligns with Go's preference for explicit code over hidden magic.",
	"The introduction of generics in Go 1.18 was a game-changer for many developers. Generic programming allows you to write functions and data structures that work with any type while maintaining type safety. This feature reduces code duplication while preserving Go's commitment to clarity and performance.",
	"Structuring Go applications properly is crucial for maintainability. The standard project layout emphasizes separation of concerns, with packages organized by functionality rather than technical layers. This approach makes it easier to understand and extend codebases as they grow.",
	"Go's defer statement is a powerful mechanism for cleanup operations. It schedules a function call to be executed immediately before the surrounding function returns. This guarantees that resources like file handles and network connections are properly closed, even if the function returns unexpectedly.",
	"Working with JSON in Go is straightforward using the encoding/json package. By adding struct tags, you can control exactly how your structs are serialized and deserialized. This makes API integration painless while maintaining Go's strong typing benefits.",
	"Embedding in Go provides a simple form of composition that's often more appropriate than inheritance. When you embed a type within a struct, the embedding struct gains access to the embedded type's methods. This creates clear, composable code without complex hierarchy problems.",
	"Rate limiting is essential for robust web services. Go's time/rate package provides an elegant implementation of the token bucket algorithm. This helps protect your services from being overwhelmed by too many requests, ensuring stability under heavy load.",
	"Go's reflection capabilities, while powerful, should be used sparingly. The reflect package allows runtime inspection of types and values, which is useful for tasks like ORM implementations and generic data handling. However, reflection comes with performance costs and reduced type safety.",
	"Middleware patterns in Go web applications provide a clean way to separate cross-cutting concerns like logging, authentication, and error handling. By implementing the http.Handler interface, middleware components can be composed to build a processing pipeline for HTTP requests.",
	"Understanding Go's slices is fundamental to effective programming in the language. Unlike arrays, slices are reference types that provide a view into an underlying array. This distinction affects how they behave when passed to functions or assigned to variables.",
	"Go modules solved dependency management issues that plagued earlier Go projects. With explicit versioning and reproducible builds, modules ensure that your code compiles the same way everywhere. The go.mod file makes dependencies explicit and version requirements clear.",
	"Channels in Go provide a safe way for goroutines to communicate. By using channels rather than shared memory, you avoid many common concurrency bugs. Remember the Go proverb: 'Don't communicate by sharing memory; share memory by communicating.'",
	"Building command-line tools in Go is straightforward with packages like cobra or flag. Go's fast compilation and single binary output make it ideal for CLI applications. Users don't need runtime environments installed - just distribute your compiled binary and they're ready to go.",
}

var blogTags = []string{
	"golang",
	"concurrency",
	"goroutines",
	"web-development",
	"microservices",
	"performance",
	"testing",
	"api",
	"databases",
	"docker",
	"kubernetes",
	"backend",
	"best-practices",
	"error-handling",
	"security",
	"json",
	"rest-api",
	"go-modules",
	"interfaces",
	"generics",
}

var blogComments = []string{
	"This tutorial saved me hours of debugging! Thanks for the clear explanation.",
	"I've been using Go for 5 years and still learned something new from this article.",
	"Have you considered covering how this would work in a distributed system?",
	"The code examples are great, but I think there's a small bug in the third snippet.",
	"How would this approach scale with larger datasets?",
	"I implemented this pattern in our production code and it reduced CPU usage by 30%!",
	"Great explanation of goroutines. I finally understand how they differ from threads.",
	"Could you do a follow-up post about error handling best practices?",
	"Your explanation of interfaces makes so much more sense than the official docs.",
	"I'm new to Go and this really helped me understand the language philosophy.",
	"The comparison with Rust was interesting. Would love to see more language comparisons.",
	"Have you benchmarked this solution against the standard library implementation?",
	"This is exactly what I needed for a project I'm working on. Perfect timing!",
	"The diagrams really helped visualize the concurrency patterns.",
	"I've bookmarked this as a reference for my team. Excellent resource!",
	"One question: how would you handle timeouts in this scenario?",
	"After reading this, I refactored our codebase and it's much cleaner now.",
	"I appreciate the focus on both performance and readability in your examples.",
	"This is now my go-to example when teaching new developers about channels.",
	"Looking forward to the next post in this series!",
}

func Seed(store store.Storage) {
	ctx := context.Background()
	users := generateUsers(100)

	for _, user := range users {
		err := store.Users.Create(ctx, user)

		if err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)

	for _, post := range posts {
		err := store.Posts.Create(ctx, post)
		if err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		err := store.Comments.Create(ctx, comment)
		if err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: gofakeit.Username(),
			Email:    gofakeit.Email(),
			Password: "123123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		log.Println(user, user.Username)

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   blogTitles[rand.Intn(len(blogTitles))],
			Content: blogContents[rand.Intn(len(blogContents))],
			Tags: []string{
				blogTags[rand.Intn(len(blogTags))],
				blogTags[rand.Intn(len(blogTags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comments := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		comments[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: blogComments[rand.Intn(len(blogComments))],
		}
	}

	return comments
}
