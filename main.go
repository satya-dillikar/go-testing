package main

import (
	"fmt"
)

func main() {

	blog := New()

	fmt.Println(blog)

	blog.SaveArticle(Article{"My First Blog post", "Today, we will be talking about blogging"})

	fmt.Println(blog)

}
