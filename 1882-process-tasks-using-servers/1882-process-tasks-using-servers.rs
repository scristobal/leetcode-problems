
use std::cmp::Ordering;
use std::collections::{BinaryHeap, VecDeque};

#[derive(Debug, Copy, Clone, Eq, PartialEq)]
struct Server {
    weight: i32,
    index: usize,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq)]
struct Task {
    scheduled: i32,
    duration: i32,
}

impl PartialOrd for Server {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for Server {
    fn cmp(&self, other: &Self) -> Ordering {
        other
            .weight
            .cmp(&self.weight)
            .then_with(|| other.index.cmp(&self.index))
    }
}

#[derive(Debug, Copy, Clone, Eq, PartialEq)]
struct Job {
    server: Server,
    ready_time: i32,
}

impl Ord for Job {
    fn cmp(&self, other: &Self) -> Ordering {
        other.ready_time.cmp(&self.ready_time).then_with(|| {
            other
                .server
                .weight
                .cmp(&self.server.weight)
                .then_with(|| other.server.index.cmp(&self.server.index))
        })
    }
}

impl PartialOrd for Job {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Solution {
    pub fn assign_tasks(servers: Vec<i32>, tasks: Vec<i32>) -> Vec<i32> {
        let mut free_servers = servers
            .iter()
            .enumerate()
            .map(|server| Server {
                weight: *server.1,
                index: server.0,
            })
            .collect::<BinaryHeap<Server>>();

        let mut busy_servers = BinaryHeap::<Job>::new();

        let mut res: Vec<usize> = vec![];

        for (start, &duration) in tasks.iter().enumerate() {
            while !busy_servers.is_empty()
                && busy_servers.peek().unwrap().ready_time <= start as i32
            {
                let job = busy_servers.pop().unwrap();
                free_servers.push(job.server);
            }

            let server = free_servers.pop();

            match server {
                Some(server) => {
                    let ready_time = start as i32 + duration;
                    busy_servers.push(Job { server, ready_time });
                    res.push(server.index);
                }
                None => {
                    let mut job = busy_servers.pop().unwrap();
                    job.ready_time += duration;
                    busy_servers.push(job);
                    res.push(job.server.index);
                }
            }
        }

        res.iter().map(|i| *i as i32).collect()
    }
}