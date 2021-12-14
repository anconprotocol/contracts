use std::collections::HashMap;
use crate::Context;
use juniper::Variables;
use crate::schema;
use wasm_bindgen::convert::IntoWasmAbi;
use wasm_bindgen::prelude::*;
use wasm_bindgen_futures::*;

#[wasm_bindgen]
extern "C" {
    #[wasm_bindgen]
    pub fn write_store(key: String, value: String);

    #[wasm_bindgen]
    pub fn read_store(key: String) -> String;

    #[wasm_bindgen]
    pub fn write_dag_block(data: String) -> String;

    #[wasm_bindgen]
    pub fn read_dag_block(cid: String, path: String) -> String; 
}
