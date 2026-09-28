// Package library моделює простий домен «Бібліотека».
//
// Завдання 1: спроєктуйте структури Author, Book, Library самостійно
// (або з допомогою ШІ) і реалізуйте метод AddBook.
//
// Завдання 2: реалізуйте функцію SortByYear, яка сортує книги
// за PublishedYear без сторонніх бібліотек (sort.Slice або ручний
// алгоритм — обидва варіанти приймаються).
package library

import (
	"sort"
)

// Author представляє автора книги.
//
// TODO (Завдання 1): додайте/скоригуйте поля на свій розсуд —
// мають бути щонайменше Name і Born.
type Author struct {
	Name string
	Born int
}

// Book представляє одну книгу в бібліотеці.
//
// Book.Author — приклад вкладеної структури, як і вимагає завдання.
//
// TODO (Завдання 1): додайте/скоригуйте поля на свій розсуд —
// мають бути щонайменше Title, Author і PublishedYear.
type Book struct {
	Title         string
	Author        Author
	PublishedYear int
}

// BookID — ідентичність книги: за ним AddBook вирішує,
// чи є вона вже в каталозі.
type BookID struct {
	Title  string
	Author string
}

// ID повертає ключ, за яким два записи вважаються однією книгою.
//
// Явний ключ замість порівняння всієї структури: інакше книга з
// друкарською помилкою в імені автора або з незаповненим Author.Born
// потрапила б до каталогу вдруге.
func (b Book) ID() BookID {
	return BookID{Title: b.Title, Author: b.Author.Name}
}

// Library зберігає колекцію книг.
//
// Library.Books — поле-зріз, як і вимагає завдання.
type Library struct {
	Name  string
	Books []Book
}

// AddBook додає книгу до бібліотеки.
//
// TODO (Завдання 1): реалізуйте додавання b до l.Books.
func (l *Library) AddBook(b Book) bool {
	id := b.ID()

	for i := range l.Books {
		if l.Books[i].ID() == id {
			return false
		}
	}

	l.Books = append(l.Books, b)

	return true
}

// SortByYear сортує books за PublishedYear (за зростанням) на місці,
// без використання сторонніх бібліотек.
//
// TODO (Завдання 2): реалізуйте сортування.
// Підказка: sort.Slice(books, func(i, j int) bool { ... }) —
// це вже частина стандартної бібліотеки Go, тому дозволена.
func SortByYear(books []Book) {
	if len(books) < 2 {
		return
	}

	sort.Slice(books, func(i, j int) bool {
		return books[i].PublishedYear < books[j].PublishedYear
	})

	// additional implementation without sort.Slice
	// quickSort(books, func(a, b Book) bool {
	//	return a.PublishedYear < b.PublishedYear
	// })
}

//func quickSort[T any](items []T, less func(a, b T) bool) {
//	if len(items) < 2 {
//		return
//	}
//
//	pivot := items[len(items)/2]
//	lt, cur, gt := 0, 0, len(items)-1
//
//	for cur <= gt {
//		switch {
//		case less(items[cur], pivot):
//			items[lt], items[cur] = items[cur], items[lt]
//			lt++
//			cur++
//		case less(pivot, items[cur]):
//			items[cur], items[gt] = items[gt], items[cur]
//			gt--
//		default:
//			cur++
//		}
//	}
//
//	quickSort(items[:lt], less)
//	quickSort(items[gt+1:], less)
//}
