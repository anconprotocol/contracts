use serde_json::json;
use std::collections::HashMap;
use crate::Context;
use juniper::Variables;
use crate::schema;
use wasm_bindgen::convert::IntoWasmAbi;
use wasm_bindgen::prelude::*;
