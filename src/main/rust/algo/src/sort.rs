use std::fmt::Debug;

// O(n^2)
fn bubble_sort<T: PartialOrd + Debug>(v: &mut [T]) {
    for p in 0..v.len() {
        let mut sorted = true;
        for i in 0..(v.len() - 1)-p {
            if v[i] > v[i + 1] {
                v.swap(i, i + 1);
                sorted = false;
            }
        }
        if sorted {
            return;
        }
    }
}

pub fn sort() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test1() {
        // let mut i: Vec<int> = Vec::new(1, 2, 3);
        let mut i: Vec<i32> = vec![1, 2, 3];
        bubble_sort(&mut i);
        assert_eq!(i, vec![1, 2, 3]);
    }
}