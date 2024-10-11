#[cfg(test)]
mod shortinette_rust_test_module05_ex01_0001 {
    use super::*;

    #[test]
    fn no_buffer_empty_message() {
        let mut out = Vec::new();
        let mut log = Logger::new(0, &mut out);

        log.log("").unwrap();
        assert_eq!(out, b"\n");
    }

    #[test]
    fn no_buffer_sinlge_letter() {
        let mut out = Vec::new();
        let mut logger = Logger::new(0, &mut out);

        logger.log("h").unwrap();
        assert_eq!(out, b"h\n");
    }

    #[test]
    fn no_buffer_sinle_word() {
        let mut out = Vec::new();
        let mut logger = Logger::new(0, &mut out);

        logger.log("hello").unwrap();
        assert_eq!(out, b"hello\n");
    }

    #[test]
    fn no_buffer_multiple_words() {
        let mut out = Vec::new();
        let mut logger = Logger::new(0, &mut out);

        logger.log("hello").unwrap();
        logger.log("world").unwrap();
        logger.log("testing").unwrap();
        logger.log("h").unwrap();
        logger.log("").unwrap();
        assert_eq!(out, b"hello\nworld\ntesting\nh\n\n");
    }

    #[test]
    fn buffered_empty_message() {
        let mut out = Vec::new();
        let mut logger = Logger::new(12, &mut out);

        logger.log("").unwrap();
        assert_eq!(logger.writer, b"");

        logger.flush().unwrap();
        assert_eq!(logger.writer, b"\n")
    }

    #[test]
    fn buffered_message() {
        let mut out = Vec::new();
        let mut logger = Logger::new(12, &mut out);

        logger.log("hello").unwrap();
        assert_eq!(logger.writer, b"");

        logger.flush().unwrap();
        assert_eq!(logger.writer, b"hello\n");
    }

    #[test]
    fn buffered_messages() {
        let mut out = Vec::new();
        let mut logger = Logger::new(12, &mut out);

        logger.log("hello").unwrap();
        assert_eq!(logger.writer, b"");

        logger.log("world").unwrap();
        assert_eq!(logger.writer, b"hello\nworld\n");
    }

    #[test]
    fn buffer_len_same_as_message() {
        let mut out = Vec::new();
        let mut logger = Logger::new(12, &mut out);

        logger.log("Hello World!").unwrap();
        assert_eq!(logger.writer, b"Hello World!\n");
    }

    #[test]
    fn buffer_len_same_as_message_with_newline() {
        let mut out = Vec::new();
        let mut logger = Logger::new(12, &mut out);

        logger.log("Hello World").unwrap();
        assert_eq!(logger.writer, b"Hello World\n");
    }

    #[test]
    fn buffer_too_long_message() {
        let mut out = Vec::new();
        let mut logger = Logger::new(1024, &mut out);

        logger.log(&"a".repeat(2048)).unwrap();
        assert_eq!(logger.writer, format!("{}\n", "a".repeat(2048)).as_bytes());
    }

    #[test]
    fn empty_buffer_flush() {
        let mut out = Vec::new();
        let mut logger = Logger::new(1024, &mut out);

        logger.flush().unwrap();
        assert!(logger.writer.is_empty());
    }

    #[test]
    fn correct_buffer_size() {
        let logger = Logger::new(0, vec![0; 0]);
        assert_eq!(0, logger.buffer.len());

        let logger = Logger::new(128, vec![0; 0]);
        assert_eq!(128, logger.buffer.len());

        let logger = Logger::new(1024, vec![0; 0]);
        assert_eq!(1024, logger.buffer.len());
    }
}
