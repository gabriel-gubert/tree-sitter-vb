import XCTest
import SwiftTreeSitter
import TreeSitterVb

final class TreeSitterVbTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_vb())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading no grammar")
    }
}
