use std::fmt::Display;
use std::collections::BTreeMap;

/* Builder Objects and Default Properties */

#[derive(Debug, PartialEq)]
pub enum Property {
    Simple(&'static str, String),
    Style(&'static str, String),
    Transform(String),
}

#[derive(Debug, PartialEq)]
pub struct SvgTag {
    pub kind: &'static str,
    pub properties: Vec<Property>,
    pub children: Vec<SvgTag>,
}

#[derive(Debug, PartialEq)]
pub enum Ability {
    Charge,
    Taunt,
}

#[derive(Debug, PartialEq, PartialOrd, Ord, Eq)]
pub enum Trigger {
    BattleCry,
    Death,
}

#[derive(Debug, PartialEq, Default)]
pub struct Card {
    pub name: String,
    pub abilities: Vec<Ability>,
    pub cost: i32,
    pub triggers: BTreeMap<Trigger, String>,
}
#[derive(Debug, PartialEq, Default)]
pub struct CardBuilder {
    pub name: String,
    pub abilities: Vec<Ability>,
    pub cost: Option<i32>,
    pub triggers: BTreeMap<Trigger, String>,
}

impl CardBuilder {
    pub fn new(name: String) -> Self {
        CardBuilder {
            name,
            cost: None,
            ..Default::default()
        }
    }
    pub fn builld(self) -> Card {
        Card {
            name: self.name,
            cost: self.cost.unwrap_or(1),
            abilities: self.abilities,
            triggers: self.triggers,
        }
    }
}

impl Card {
    pub fn build(name: String) -> CardBuilder {
        CardBuilder::new(name)
    }
}

impl SvgTag {
    pub fn new(kind: &'static str) -> Self {
        SvgTag {
            kind,
            properties: Vec::new(),
            children: Vec::new(),
        }
    }
    pub fn child(mut self, c: SvgTag) -> Self {
        self.children.push(c);
        self
    }
    pub fn property<V: Display>(mut self, k: &'static str, v: V) -> Self {
        self.properties.push(Property::Simple(k, v.to_string()));
        self
    }
    pub fn style<V: Display>(mut self, k: &'static str, v: V) -> Self {
        self.properties.push(Property::Style(k, v.to_string()));
        self
    }
    pub fn x<V: Display>(self, v: V) -> Self {
        self.property("x", v)
    }
    pub fn y<V: Display>(self, v: V) -> Self {
        self.property("y", v)
    }
    pub fn w<V: Display>(self, v: V) -> Self {
        self.property("width", v)
    }
    pub fn h<V: Display>(self, v: V) -> Self {
        self.property("height", v)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::Property::Simple;

    #[test]
    fn builderTest() {
        // builder pattern
        let a = SvgTag::new("rect").x("5").y("5").w("50").h("20");
        let b = SvgTag::new("svg").w("60px").h("80px").child(a);
    }
    #[test]
    fn cardBuilderTest() {
        let x = Card::build("General".to_string()).builld();
    }
}

fn main() -> Result<(), std::io::Error> {
    Ok(())
}