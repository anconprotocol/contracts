use crate::sdk::{read_dag_block, write_dag_block};

#[macro_use]
extern crate juniper_codegen;

pub mod sdk;
pub mod contract;

use crate::contract::{Context, schema, Ancon721Metadata};
use wasm_bindgen::prelude::*;
extern crate juniper;

use juniper::{
    Variables,
};
use serde_json::json;

use std::collections::HashMap;

use std::str;
use std::vec::*;

#[wasm_bindgen()]
pub fn execute(query: &str) -> String {
    let ctx = Context {
        metadata: HashMap::default(),
    };

    let v = Variables::new();
    let sch = schema();
    let res = juniper::execute_sync(query, None, &sch, &v, &ctx);
    let (data, err) = res.unwrap();
    let errors = err
        .iter()
        .map(|i| i.error().message().to_string())
        .collect::<Vec<String>>();

    json!({
        "data":data.to_string(),
        "errors": errors,
    })
    .to_string()
}

#[wasm_bindgen]
pub fn store(data: &str) -> Vec<u8> {
    let payload = Ancon721Metadata {
        name: "test".to_string(),
        description: "description".to_string(),
        image: "http://ipfs.io/ipfs/".to_string(),
        owner: "".to_string(),
        parent: "".to_string(),
        sources: [].to_vec(),
    };

    let json_payload = serde_json::to_string(&payload).unwrap();
    unsafe {
        let l = json_payload.len() as usize;
        let metadata = write_dag_block(&json_payload).to_vec();
        metadata[..l].to_vec()
    }
}
