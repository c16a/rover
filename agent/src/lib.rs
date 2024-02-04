use warp::Filter;
pub fn add(left: usize, right: usize) -> usize {
    left + right
}

pub async fn start() {
    let hello = warp::path!("hello" / String).map(|name| format!("Hello, {}!", name));

    println!("Starting Rover agent");

    warp::serve(hello).run(([127, 0, 0, 1], 3030)).await;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
