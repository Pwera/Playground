use std::ops::Add;
use std::ops::AddAssign;

mod combi;
mod itertools;

pub trait Rangable: AddAssign + PartialOrd + Copy {}

impl<T: AddAssign + PartialOrd + Copy> Rangable for T {}

pub struct RangeIterator<T> {
    curr: T,
    step: T,
    stop: T,
}

impl<T: Rangable> RangeIterator<T> {
    pub fn new(start: T, stop: T, step: T) -> Self {
        // SkipIterator{};
        RangeIterator {
            curr: start,
            stop,
            step,
        }
    }
}

impl<T: Rangable> Iterator for RangeIterator<T> {
    type Item = T;

    fn next(&mut self) -> Option<Self::Item> {
        if self.curr >= self.stop {
            return None;
        }
        let res = self.curr;
        self.curr += self.step;
        Some(res)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::combi::IterCombi;

    #[test]
    fn test_range_iterator_new_for_integers() {
        let mut m = 0;
        let it = RangeIterator::new(5, 12, 3);
        for s in it {
            m += s;
        }
        assert_eq!(m, 5 + 8 + 11);
    }

    #[test]
    fn test_range_iterator_new_for_float() {
        let mut m = 0.;
        let it = RangeIterator::new(5., 12., 2.5);
        for s in it {
            m += s;
        }
        assert_eq!(m, 5. + 7.5 + 10.);
    }

    #[test]
    fn test_range_iterator_filter() {
        let v: i32 = RangeIterator::new(3, 13, 3)
            .filter(|x| x % 2 == 0)
            .sum();
        assert_eq!(v, 6 + 12);
    }

    #[test]
    fn test_range_iterator_skiphalf() {
        let _v: i32 = RangeIterator::new(3, 13, 3)
            .skip_half()
            .sum();
        let v2: i32 = (0..10).skip_half().sum();
        assert_eq!(v2, 1 + 3 + 5 + 7 + 9);
    }
}
