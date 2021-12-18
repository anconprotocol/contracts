
pub fn focused_transform_patch_str(cid: &str, path: &str, prev: &str, next: &str) ->  [u8; 1024]{

}

pub fn read_dag(cid: &str) -> Vec<u8> {
    unsafe {
        let s: i32 = 0;
        let metadata = read_dag_block(&cid, &s);
        let x = s as usize;
        let v = &metadata[..x];
        v.to_vec()
    }
}

extern "C" {
    #[no_mangle]
    pub fn write_store(key: &str) -> [u8;1024];

    #[no_mangle]
    pub fn read_store(key: &str) -> [u8;1024];

    #[no_mangle]
    pub fn write_dag_block(data: &str) -> [u8;1024];

    #[no_mangle]
    pub fn read_dag_block(cid: &str, ret: &i32) -> [u8;1024];
}
