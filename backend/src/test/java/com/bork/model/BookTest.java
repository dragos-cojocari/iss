package com.bork.model;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

class BookTest {

    private Book book;
    private Category category;

    @BeforeEach
    void setUp() {
        category = new Category();
        category.setCategoryId(UUID.randomUUID());
        category.setName("Science Fiction");

        book = new Book();
        book.setBookId(UUID.randomUUID());
        book.setTitle("The Matrix");
        book.setAuthor("Wachowski Sisters");
        book.setIsbn("978-0-123456-78-9");
        book.setCategory(category);
        book.setIsAvailable(true);
    }

    @Test
    void bookCreation_ShouldHaveDefaultValues() {
        Book newBook = new Book();

        assertNotNull(newBook);
        assertTrue(newBook.getIsAvailable());
    }

    @Test
    void bookSettersAndGetters_ShouldWorkCorrectly() {
        UUID bookId = UUID.randomUUID();
        book.setBookId(bookId);
        book.setTitle("New Title");
        book.setAuthor("New Author");
        book.setIsbn("978-1-234567-89-0");
        book.setIsAvailable(false);

        assertEquals(bookId, book.getBookId());
        assertEquals("New Title", book.getTitle());
        assertEquals("New Author", book.getAuthor());
        assertEquals("978-1-234567-89-0", book.getIsbn());
        assertFalse(book.getIsAvailable());
    }

    @Test
    void bookCategory_ShouldBeAssignable() {
        Category newCategory = new Category();
        newCategory.setCategoryId(UUID.randomUUID());
        newCategory.setName("Fantasy");

        book.setCategory(newCategory);

        assertEquals(newCategory, book.getCategory());
        assertEquals("Fantasy", book.getCategory().getName());
    }

    @Test
    void bookAvailability_ShouldBeToggleable() {
        assertTrue(book.getIsAvailable());

        book.setIsAvailable(false);
        assertFalse(book.getIsAvailable());

        book.setIsAvailable(true);
        assertTrue(book.getIsAvailable());
    }
}
