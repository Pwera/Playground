use std::sync::Mutex;

lazy_static::lazy_static! {
    static ref RG: Mutex<RandGen> = Mutex::new(RandGen::new(12));
}

pub fn rand(max: usize) -> usize {
    RG.lock().unwrap().next_v(max)
}


struct RandGen {
    curr: usize,
    mul: usize,
    inc: usize,
    _modulo: usize,
}

impl RandGen {
    fn new(curr: usize) -> Self {
        RandGen {
            curr,
            mul: 56394237,
            inc: 346423491,
            _modulo: 23254544561,
        }
    }

    fn next_v(&mut self, max: usize) -> usize {
        self.curr = (self.curr * self.mul + self.inc) % self._modulo;
        self.curr % max
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test1() {
        let mut gen = RandGen::new(12);
        for _ in 0..100 {
            println!("-> {}", gen.next_v(100));
        }
    }
}