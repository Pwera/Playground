pub struct Company {
    ceo: String,
    receptionist: String,
    marketing: String,
}

pub struct CompanyIter<'a> {
    c: &'a Company,
    n: i32,
}

impl<'a> Iterator for CompanyIter<'a> {
    type Item = &'a str;

    fn next(&mut self) -> Option<Self::Item> {
        self.n += 1;
        match self.n {
            1 => Some(&self.c.ceo),
            2 => Some(&self.c.receptionist),
            3 => Some(&self.c.marketing),
            _ => None,
        }
    }
}

impl<'a> IntoIterator for &'a Company {
    type Item = &'a str;
    type IntoIter = CompanyIter<'a>;

    fn into_iter(self) -> Self::IntoIter {
        CompanyIter { c: self, n: 0 }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_company_() {
        let c = Company {
            ceo: "Alice".to_string(),
            receptionist: "Bob".to_string(),
            marketing: "Chad".to_string(),
        };

        let mut res = String::new();

        // let v= vec![c];
        for m in &c {
            res.push_str(m);
        }
        assert_eq!(res, "AliceBobChad");
    }

    #[test]
    fn test_itertools_step_by() {
        let v: i32 = (0..10).step_by(3).sum();
        assert_eq!(v, 0 + 3 + 6 + 9);
    }

    use itertools::Itertools;
    use crate::itertools::Company;


    #[test]
    fn test_itertools_interleave() {
        let v: Vec<i32> = (0..4)
            .interleave(
                (11..15).rev()
            ).collect();
        assert_eq!(v, vec![0, 14, 1, 13, 2, 12, 3, 11]);
    }

    #[test]
    fn test_itertools_intersperse() {
        let s = "hello world etc";
        let v: Vec<&str> = s.split(" ").intersperse(",").collect();
        assert_eq!(v, vec!["hello", ",", "world", ",", "etc"]);
        println!("asc");
    }
}