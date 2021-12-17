    
extern "C" {
    #[no_mangle]
    pub fn write_store(key: &str, value: &str) -> [u8;1024];

    #[no_mangle]
    pub fn read_store(key: &str) -> [u8;1024];

    #[no_mangle]
    pub fn write_dag_block(data: &str) -> [u8;1024];

    #[no_mangle]
    pub fn read_dag_block(cid: &str, path: &str) -> [u8;1024];
}
