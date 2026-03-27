package main

import "fmt"

func main() {
    for {
        var input1 int
        var input2 int
        var operator string

        fmt.Print("enter an input: ")
        fmt.Scan(&input1)

        fmt.Print("enter an input: ")
        fmt.Scan(&input2)

        fmt.Println("<+> addition")
        fmt.Println("<-> subtraction")
        fmt.Println("<*> multiplication")
        fmt.Println("</> division")
        fmt.Println("<quit> exit program")
        fmt.Println("<help> ask for help")

        fmt.Println("enter an operator: ")
        fmt.Scan(&operator)

        if operator == "quit" {
            fmt.Println("good bye")
            break
        }

        if operator == "help" {
            fmt.Println("Commands:")
            fmt.Println("add a b - addition")
            fmt.Println("sub a b - subtraction")
            fmt.Println("mul a b - multiplication")
            fmt.Println("div a b - division")
            fmt.Println("quit - exit program")
            continue
        }

        switch operator {
        case "+":
            fmt.Println("result =", input1+input2)

        case "-":
            fmt.Println("result =", input1-input2)

        case "*":
            fmt.Println("result =", input1*input2)

        case "/":
            if input2 == 0 {
                fmt.Println("invalid division")
            } else {
                fmt.Println("result =", input1/input2)
            }
        default:
            fmt.Println("invalid input")
            continue

        }
        if operator == "quit" {
            fmt.Println("good bye")
            break
        }

        if operator == "help" {
            fmt.Println("Commands:")
            fmt.Println("add a b - addition")
            fmt.Println("sub a b - subtraction")
            fmt.Println("mul a b - multiplication")
            fmt.Println("div a b - division")
            fmt.Println("quit - exit program")
            continue

        }

    }
}

