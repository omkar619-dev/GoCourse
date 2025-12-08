package main

func main() {
 fruit := "apple"
 switch fruit {
 case "banana":
	 println("This is a banana.")
 case "apple":
	 println("This is an apple.")
	 fallthrough
 case "orange":
	 println("This is an orange.")
 default:
	 println("Unknown fruit.")
 }
}