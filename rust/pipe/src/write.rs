use crossbeam::channel::Receiver;
use std::fs::File;
use std::io::{self, BufWriter, ErrorKind, Result, Write};

pub fn write_loop(outfile: &str, receiver: Receiver<Vec<u8>>) -> Result<bool> {
    let mut writer: Box<dyn Write> = if !outfile.is_empty() {
        Box::new(BufWriter::new(File::create(outfile)?))
    } else {
        Box::new(io::stdout())
    };

    loop {
        let result = receiver.recv().unwrap();
        if result.is_empty() {
            break;
        }
        if let Err(e) = writer.write_all(&result) {
            if e.kind() == ErrorKind::BrokenPipe {
                return Ok(false);
            }
            return Err(e);
        }
    }
    Ok(true)
}
