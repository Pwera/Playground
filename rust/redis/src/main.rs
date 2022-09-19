use std::error::Error;
use redis::{from_redis_value, streams::{StreamRangeReply, StreamReadOptions, StreamReadReply},
            AsyncCommands, Client};
use std::thread::sleep;
use tokio::time::Duration;

#[tokio::main(flavor="current_thread")]
async fn main() -> redis::RedisResult<()> {
    let my_stream_name = "my_stream";
    let client = redis::Client::open("redis://127.0.0.1/").unwrap();
    let mut con = client.get_async_connection().await?;

    con.set("my_key", "Ok").await?;
    let res: String = con.get("my_key").await?;
    println!("res: {}", res);

    /* Redis Streams*/

    con.xadd(my_stream_name, "*", &[("name", "01"), ("name", "02")]).await?;
    let len: i32 = con.xlen(my_stream_name).await?;
    println!("stream length: {}", len);


    /* xrevrange the read stream */
    let result: Option<StreamRangeReply> = con.xrevrange_count(my_stream_name, "+", "-", 10).await?;
    if let Some(reply) = result {
        for stream_id in reply.ids {
            println!("xrevrange: {}", stream_id.id);
            for (name, value) in stream_id.map.iter() {
                println!("{} -> {}", name, from_redis_value::<String>(value)?);
            }
            println!("");
        }
    }

    /* Blocking xread */
    tokio::spawn(async {
        let my_stream_name = "my_stream";
        let client = redis::Client::open("redis://127.0.0.1/").unwrap();
        let mut con = client.get_async_connection().await.unwrap();
        loop {
            let options = StreamReadOptions::default().count(1).block(0);
            let result: Option<StreamReadReply> = con.xread_options(&[my_stream_name], &["$"], &options).await.unwrap();
            if let Some(reply) = result {
                reply.keys.iter().for_each(|stream_key| {
                    println!("xread block: {}", stream_key.key);
                    stream_key.ids.iter().for_each(|stream_id| { println!("stream id: {:?}", stream_id) });
                });
            }
        }
    });

    sleep(Duration::from_millis(1000));
    con.xadd(my_stream_name, "*", &[("name", "01"), ("name", "02")]).await?;

    println!("done");
    con.del("my_key").await?;
    con.del(my_stream_name).await?;
    Ok(())
}
