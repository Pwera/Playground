//!  # Pipe library
pub mod args;
pub mod read;
pub mod stats;
pub mod write;

/// CHUNK_SIZE constant
const CHUNK_SIZE: usize = 16 * 1024;
