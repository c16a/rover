extern crate clap;
extern crate agent;

use clap::{Args, Parser, Subcommand};
use agent::start;

#[derive(Parser)]
#[command(author, version, about, long_about = None)]
struct RoverCli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    /// Starts a rover agent
    Agent(AgentArgs)
}

#[derive(Args)]
struct AgentArgs {
    /// Path to Rover JSON config file
    #[arg(short, long)]
    config: Option<String>
}

#[tokio::main]
async fn main() {
    let cli = RoverCli::parse();

    match &cli.command {
        Commands::Agent(args) => {
            start().await;
        }
    }
}