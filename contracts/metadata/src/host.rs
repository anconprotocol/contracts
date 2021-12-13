use wasm_bindgen::convert::IntoWasmAbi;
use wasm_bindgen::prelude::*;
use wasm_bindgen_futures::*;
use std::fmt::Display;
use std::future::*;
use std::io::Cursor;

#![allow(unused)]
fn main() {
    #[wasm_bindgen]
    pub fn execute_contract(gql_query: String) -> String {
        execute(query)
    }


extern {
    #[wasm_bindgen]
    pub fn read_dag_block(cid: &str,path: &str) -> String;
}
}