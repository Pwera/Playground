// pub enum Res<T, E> {
//     Thing(T),
//     Error(E),
// }
// -> Result


pub fn devide(a: i32) -> Result<i32, String> {
    if a == 0 {
        return Result::Err("a cannot be zero".to_string());
    }
    return Result::Ok(10);
}