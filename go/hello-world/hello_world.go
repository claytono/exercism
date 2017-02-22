package hello

const testVersion = 2

func HelloWorld(name string) string {
        if name == "" {
                return "Hello, World!"
        }
        return "Hello, " +name + "!"
}
