use std::sync::mpsc::channel;
use std::thread;
use std::sync::Arc;

struct Item {
    id: i32
}

pub fn f() {
    let (sender, receiver) = channel::<Arc<Item>>();
    let sender1 = sender.clone();
    let handle = thread::Builder::new()
        .name(String::from("Thread 1"))
        .spawn(move || {
            println!("Thread: {} ", thread::current().name().unwrap());
            for  val in (1..10) {
                sender.send(Arc::new(Item { id: val }));
            }
        })
        .unwrap();

    let handle2 = thread::Builder::new()
        .name(String::from("Thread 2"))
        .spawn(move || {
            println!("Thread: {} ", thread::current().name().unwrap());
            for  val in (1..10) {
                sender1.send(Arc::new(Item { id: val }));
            }
        })
        .unwrap();

    for  val in (1..19) {
        let result = receiver.recv().unwrap().id;
        println!("{} ", result);
    }


    handle.join();
    handle2.join();
}