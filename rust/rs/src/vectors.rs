extern crate itertools;
use itertools::Itertools;

pub fn vector() {
    // Vector
    let mut a = Vec::new();
    a.push(1);
    a.push(2);
    a.push(3);
    a.push(4);
    println!("{:?}", a);
    println!("{}", a[0]);
    a[0] = 213123;
    println!("{:?}", a);
    //println!("{}", a[23]);
    //return option type
    match a.get(6) {
        Some(z) => println!("{}", z),
        None => println!("{:?}", a)
    }

    let mut c = Vec::new();
    c.push(7);

    println!("{:?}", c.pop());
    println!("{:?}", c.pop());

    while let Some(x) = a.pop() {
        println!("{} ", x);
    }
}
pub fn iterators(){
    let cities = ["Toronto", "New York", "Melbourne"];
    let populations = [2_615_060, 8_550_405, ‎4_529_500];

    let matrix = cities.iter()
        .zip(populations.iter());

    for (c, p) in matrix {
        println!("{:10}: population = {}", c, p);
    }

    let v = (1..)
        .map(|x| x * x)
        .filter(|x| x % 5 == 0 )
        .take(10)
        .collect::<Vec<i32>>();

    println!("{:?} ", v);

    for i in (0..11).step_by(2) {
        print!("{} ", i);
    }

    let data = vec![1, 4, 3, 1, 4, 2, 5];
    let unique = data.iter().unique();

    for d in unique {
        print!("{} ", d);
    }

    let creatures = vec!["banshee", "basilisk", "centaur"];
    let list = creatures.iter().join(", ");
    println!("In the enchanted forest, we found {}.", list);

    let happiness_index = vec![
        ("Canada", 7), ("Iceland", 4), ("Netherlands", 6),
        ("Finland", 1), ("New Zealand", 8), ("Denmark", 3),
        ("Norway", 2), ("Sweden", 9), ("Switzerland", 5)
    ];

    let top_contries = happiness_index
        .into_iter()
        .sorted_by(|a, b| (&a.1).cmp(&b.1))
        .into_iter()
        .take(5);

    for (country, rating) in top_contries {
        println!("# {}: {}", rating, country);
    }

    let mut vec = vec![1, 2, 3, 4, 5];
    vec.iter_mut().for_each(|el| *el *= 2);
    println!("{:?}", vec);

}