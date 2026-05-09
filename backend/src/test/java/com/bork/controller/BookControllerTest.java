package com.bork.controller;

import com.bork.model.Book;
import com.bork.model.Category;
import com.bork.repository.BookRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.ResponseEntity;

import java.util.Arrays;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class BookControllerTest {

    @Mock
    private BookRepository bookRepository;

    @InjectMocks
    private BookController bookController;

    private Book book1;
    private Book book2;
    private Category category;

    @BeforeEach
    void setUp() {
        category = new Category();
        category.setCategoryId(UUID.randomUUID());
        category.setName("Science Fiction");

        book1 = new Book();
        book1.setBookId(UUID.randomUUID());
        book1.setTitle("The Matrix");
        book1.setAuthor("Wachowski Sisters");
        book1.setIsbn("978-0-123456-78-9");
        book1.setCategory(category);
        book1.setIsAvailable(true);

        book2 = new Book();
        book2.setBookId(UUID.randomUUID());
        book2.setTitle("Neuromancer");
        book2.setAuthor("William Gibson");
        book2.setIsbn("978-0-987654-32-1");
        book2.setCategory(category);
        book2.setIsAvailable(false);
    }

    @Test
    void getAllBooks_ShouldReturnListOfBooks() {
        List<Book> books = Arrays.asList(book1, book2);
        when(bookRepository.findAll()).thenReturn(books);

        ResponseEntity<List<Book>> response = bookController.getAllBooks();

        assertNotNull(response);
        assertEquals(200, response.getStatusCode().value());
        assertNotNull(response.getBody());
        assertEquals(2, response.getBody().size());
        assertEquals("The Matrix", response.getBody().get(0).getTitle());
    }

    @Test
    void getAvailableBooks_ShouldReturnOnlyAvailableBooks() {
        List<Book> availableBooks = Arrays.asList(book1);
        when(bookRepository.findByIsAvailableTrue()).thenReturn(availableBooks);

        ResponseEntity<List<Book>> response = bookController.getAvailableBooks();

        assertNotNull(response);
        assertEquals(200, response.getStatusCode().value());
        assertNotNull(response.getBody());
        assertEquals(1, response.getBody().size());
        assertTrue(response.getBody().get(0).getIsAvailable());
    }

    @Test
    void getBookById_WithExistingId_ShouldReturnBook() {
        UUID bookId = book1.getBookId();
        when(bookRepository.findById(bookId)).thenReturn(Optional.of(book1));

        ResponseEntity<Book> response = bookController.getBookById(bookId);

        assertNotNull(response);
        assertEquals(200, response.getStatusCode().value());
        assertNotNull(response.getBody());
        assertEquals("The Matrix", response.getBody().getTitle());
    }

    @Test
    void getBookById_WithNonExistingId_ShouldReturnNotFound() {
        UUID bookId = UUID.randomUUID();
        when(bookRepository.findById(bookId)).thenReturn(Optional.empty());

        ResponseEntity<Book> response = bookController.getBookById(bookId);

        assertEquals(404, response.getStatusCode().value());
    }

    @Test
    void searchBooks_ShouldReturnMatchingBooks() {
        List<Book> searchResults = Arrays.asList(book1);
        when(bookRepository.searchByTitleOrAuthor(anyString())).thenReturn(searchResults);

        ResponseEntity<List<Book>> response = bookController.searchBooks("Matrix");

        assertNotNull(response);
        assertEquals(200, response.getStatusCode().value());
        assertNotNull(response.getBody());
        assertEquals(1, response.getBody().size());
    }
}
