#[cfg(test)]
use mockiato::mockable;

#[cfg_attr(test, mockable)]
trait Greeter {
    fn greet(&self, name: &str) -> String;
}

struct MyGreeter {}

impl Greeter for MyGreeter {
    fn greet(&self, name: &str) -> String {
        return "-> ".to_string() + name;
    }
}


struct MyService {
    pub name: String,
    pub component: Box<Greeter>,
}

impl MyService {
    pub fn new(component: Box<Greeter>) -> Self {
        Self {
            name: "".to_string(),
            component,
        }
    }
    pub fn greet(self) -> String {
        self.component.greet("MyService::greet")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test1() {
        let greeter = MyGreeter {};
        let service = MyService::new(Box::new(greeter));
        println!("{}", service.greet());
    }

    #[test]
    fn test2() {
        let mut greeter = GreeterMock::new();
        greeter
            .expect_greet(|arg| arg.partial_eq(format!("MyService::{}", "greet")))
            .times(1..2)
            .returns(String::from("MyService::mock"));
        let service = MyService::new(Box::new(greeter));
        assert_eq!("MyService::mock", service.greet());
    }
}