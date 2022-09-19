cargo fmt
exec cargo clippy
#yes | cargo run | head -n 1000000000 > /dev/null

sh -c cargo test
#sh -c cargo doc --no-deps

# vim .git/hooks/pre-commit
# chmod a+x .git/hooks/pre-commit
