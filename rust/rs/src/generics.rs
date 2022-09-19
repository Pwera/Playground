
pub fn generic() {
    let a: PointGeneric<i32> = PointGeneric { x: 0, y: 0 };
    let b: PointGeneric<f64> = PointGeneric { x: 0.0, y: 0.0 };
}

struct PointGeneric<T> {
    x: T,
    y: T,
}