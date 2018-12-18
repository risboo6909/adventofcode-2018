use shared::parse_input;
use std::cell::Cell;
use std::str::FromStr;

#[derive(Debug, PartialEq)]
struct Rectangle {
    id: usize,
    x: usize,
    y: usize,
    width: usize,
    height: usize,
    overlaps: Cell<bool>,
}

impl Rectangle {
    pub(crate) fn new(id: usize, x: usize, y: usize, width: usize, height: usize) -> Self {
        Rectangle {
            id,
            x,
            y,
            width,
            height,
            overlaps: Cell::new(false),
        }
    }

    pub(crate) fn get_st_x(&self) -> usize {
        self.x
    }

    pub(crate) fn get_end_x(&self) -> usize {
        self.x + self.width
    }

    pub(crate) fn get_st_y(&self) -> usize {
        self.y
    }

    pub(crate) fn get_end_y(&self) -> usize {
        self.y + self.height
    }

    pub(crate) fn contains(&self, x: usize, y: usize) -> bool {
        if x >= self.get_st_x()
            && x < self.get_end_x()
            && y >= self.get_st_y()
            && y < self.get_end_y()
        {
            return true;
        }
        false
    }
}

fn parse_line(input: &str) -> Rectangle {
    let tmp: Vec<&str> = input.split('@').collect();

    let id = usize::from_str(tmp[0].trim_matches(|c| c == ' ' || c == '#')).unwrap();
    let descr: Vec<&str> = tmp[1].trim().split(':').map(|item| item.trim()).collect();

    let coords: Vec<usize> = descr[0]
        .split(',')
        .map(|item| usize::from_str(item).unwrap())
        .collect();

    let sizes: Vec<usize> = descr[1]
        .split('x')
        .map(|item| usize::from_str(item).unwrap())
        .collect();

    Rectangle::new(id, coords[0], coords[1], sizes[0], sizes[1])
}

#[inline]
fn intersect(rect1: &Rectangle, rect2: &Rectangle, fiber: &mut [[u8; 1000]; 1000]) {
    for x in rect1.get_st_x()..rect1.get_end_x() {
        for y in rect1.get_st_y()..rect1.get_end_y() {
            if rect2.contains(x, y) {
                fiber[x][y] = 1;
                rect1.overlaps.set(true);
                rect2.overlaps.set(true);
            }
        }
    }
}

fn main() {
    let mut fiber = [[0u8; 1000]; 1000];
    let rectangles = parse_input(|line| parse_line(line));

    for (last_idx, rect1) in rectangles.iter().enumerate() {
        for rect2 in rectangles.iter().skip(last_idx + 1) {
            intersect(&rect1, &rect2, &mut fiber);
        }
    }

    let mut total_area = 0;

    for row in fiber.iter() {
        for cell in row.iter() {
            if *cell == 1 {
                total_area += 1;
            }
        }
    }

    println!("Total shared area: {}", total_area);

    for rect in &rectangles {
        if !rect.overlaps.get() {
            println!("Not overlapping rectangle's id: {}", rect.id);
            break;
        }
    }
}
