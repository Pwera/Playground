use futures::future::Future;
use futures::stream::Stream;
use futures::io::AsyncRead;
use futures::task::{Context, Poll};
use std::pin::Pin;


pub struct SimpleFuture {
    n: i32,
}

impl Future for SimpleFuture {
    type Output = i32;

    fn poll(self: Pin<&mut Self>, _cx: &mut Context) -> Poll<Self::Output> {
        Poll::Ready(self.n)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use futures::executor::block_on;
    use futures::future::FutureExt;
    use futures::channel::oneshot;
    use futures::{SinkExt, TryFutureExt};
    use itertools::Itertools;

    #[test]
    fn test_future_returns_a_value() {
        let f = SimpleFuture { n: 10 };
        let v = block_on(f);
        assert_eq!(v, 10);
    }

    #[test]
    fn test_future_returns_a_value_and_map() {
        let f = SimpleFuture { n: 10 };
        let v = block_on(f.map(|x| { x + 1 }));
        assert_eq!(v, 11);
    }

    #[test]
    fn test_future_returns_a_value_from_one_shot_channel() {
        let f = SimpleFuture { n: 10 };
        let (ch_s, ch_r) = oneshot::channel();
        block_on(f.map(move |x| ch_s.send(x + 5)));
        let result = block_on(ch_r).unwrap();
        assert_eq!(result, 15);
    }

    #[test]
    fn test_simpleexec() {
        let f = simpleexec(10);
        let (ch_s, ch_r) = oneshot::channel();
        block_on(f.map(move |x| ch_s.send(x + 15)));
        let result = block_on(ch_r).unwrap();
        assert_eq!(result, 35);
    }

    #[test]
    fn test_async_send() {
        let (ch_s, ch_r) = oneshot::channel();
        block_on(async move {
            let v = simpleexec(10).await;
            ch_s.send(v);
        });

        let fin = block_on(async move {
            let res = ch_r.await.unwrap();
            res + 5
        });
        assert_eq!(fin, 25);
    }
}

pub async fn simpleexec(p: i32) -> i32 {
    return p + 10;
}


fn main() {}