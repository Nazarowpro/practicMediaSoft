package main

import (
	"fmt"
)

// ==================== СТЕК (Stack) ====================
type Stack struct {
	items []interface{}
}

// NewStack создает новый стек
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push добавляет элемент на вершину стека
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop удаляет и возвращает элемент с вершины стека
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("стек пуст")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Peek возвращает элемент с вершины стека без удаления
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("стек пуст")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size возвращает размер стека
func (s *Stack) Size() int {
	return len(s.items)
}

// Print выводит содержимое стека
func (s *Stack) Print() {
	fmt.Println("Стек (сверху вниз):")
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Printf("  %v\n", s.items[i])
	}
}

// ==================== ОЧЕРЕДЬ (Queue) ====================
type Queue struct {
	items []interface{}
}

// NewQueue создает новую очередь
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Enqueue добавляет элемент в конец очереди
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue удаляет и возвращает элемент из начала очереди
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("очередь пуста")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Front возвращает первый элемент очереди без удаления
func (q *Queue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("очередь пуста")
	}
	return q.items[0], nil
}

// IsEmpty проверяет, пуста ли очередь
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size возвращает размер очереди
func (q *Queue) Size() int {
	return len(q.items)
}

// Print выводит содержимое очереди
func (q *Queue) Print() {
	fmt.Println("Очередь (спереди назад):")
	for i, item := range q.items {
		fmt.Printf("  [%d] %v\n", i, item)
	}
}

// ==================== БИНАРНОЕ ДЕРЕВО (Binary Tree) ====================

// TreeNode - узел бинарного дерева
type TreeNode struct {
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree - бинарное дерево
type BinaryTree struct {
	Root *TreeNode
}

// NewBinaryTree создает новое бинарное дерево
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

// Insert добавляет значение в дерево (для целых чисел)
func (bt *BinaryTree) Insert(value int) {
	newNode := &TreeNode{Value: value, Left: nil, Right: nil}
	
	if bt.Root == nil {
		bt.Root = newNode
		return
	}
	
	bt.insertNode(bt.Root, newNode)
}

// insertNode рекурсивно вставляет узел в дерево (для целых чисел)
func (bt *BinaryTree) insertNode(node, newNode *TreeNode) {
	// Предполагаем, что значения - целые числа
	nodeVal := node.Value.(int)
	newVal := newNode.Value.(int)
	
	if newVal < nodeVal {
		// Вставляем в левое поддерево
		if node.Left == nil {
			node.Left = newNode
		} else {
			bt.insertNode(node.Left, newNode)
		}
	} else {
		// Вставляем в правое поддерево
		if node.Right == nil {
			node.Right = newNode
		} else {
			bt.insertNode(node.Right, newNode)
		}
	}
}

// Search ищет значение в дереве (для целых чисел)
func (bt *BinaryTree) Search(value int) bool {
	return bt.searchNode(bt.Root, value)
}

func (bt *BinaryTree) searchNode(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	
	nodeVal := node.Value.(int)
	
	if value == nodeVal {
		return true
	} else if value < nodeVal {
		return bt.searchNode(node.Left, value)
	} else {
		return bt.searchNode(node.Right, value)
	}
}

// InOrder обходит дерево в симметричном порядке (лево-корень-право)
func (bt *BinaryTree) InOrder() {
	fmt.Print("In-order обход: ")
	bt.inOrderTraversal(bt.Root)
	fmt.Println()
}

func (bt *BinaryTree) inOrderTraversal(node *TreeNode) {
	if node != nil {
		bt.inOrderTraversal(node.Left)
		fmt.Printf("%v ", node.Value)
		bt.inOrderTraversal(node.Right)
	}
}

// PreOrder обходит дерево в прямом порядке (корень-лево-право)
func (bt *BinaryTree) PreOrder() {
	fmt.Print("Pre-order обход: ")
	bt.preOrderTraversal(bt.Root)
	fmt.Println()
}

func (bt *BinaryTree) preOrderTraversal(node *TreeNode) {
	if node != nil {
		fmt.Printf("%v ", node.Value)
		bt.preOrderTraversal(node.Left)
		bt.preOrderTraversal(node.Right)
	}
}

// PostOrder обходит дерево в обратном порядке (лево-право-корень)
func (bt *BinaryTree) PostOrder() {
	fmt.Print("Post-order обход: ")
	bt.postOrderTraversal(bt.Root)
	fmt.Println()
}

func (bt *BinaryTree) postOrderTraversal(node *TreeNode) {
	if node != nil {
		bt.postOrderTraversal(node.Left)
		bt.postOrderTraversal(node.Right)
		fmt.Printf("%v ", node.Value)
	}
}

// ==================== ТЕСТИРОВАНИЕ ====================
func main() {
	fmt.Println("=== ТЕСТИРОВАНИЕ СТРУКТУР ДАННЫХ ===\n")

	// Тестирование стека
	fmt.Println("--- СТЕК ---")
	stack := NewStack()
	
	// Добавляем элементы
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push("Hello")
	stack.Print()
	
	fmt.Printf("Размер стека: %d\n", stack.Size())
	
	// Извлекаем элементы
	item, _ := stack.Pop()
	fmt.Printf("Извлечен элемент: %v\n", item)
	item, _ = stack.Pop()
	fmt.Printf("Извлечен элемент: %v\n", item)
	
	// Смотрим вершину
	peek, _ := stack.Peek()
	fmt.Printf("Вершина стека: %v\n", peek)
	
	stack.Print()
	fmt.Println()

	// Тестирование очереди
	fmt.Println("--- ОЧЕРЕДЬ ---")
	queue := NewQueue()
	
	// Добавляем элементы
	queue.Enqueue("Первый")
	queue.Enqueue("Второй")
	queue.Enqueue("Третий")
	queue.Enqueue(42)
	queue.Print()
	
	fmt.Printf("Размер очереди: %d\n", queue.Size())
	
	// Извлекаем элементы
	front, _ := queue.Front()
	fmt.Printf("Первый элемент: %v\n", front)
	
	item, _ = queue.Dequeue()
	fmt.Printf("Извлечен элемент: %v\n", item)
	item, _ = queue.Dequeue()
	fmt.Printf("Извлечен элемент: %v\n", item)
	
	queue.Print()
	fmt.Println()

	// Тестирование бинарного дерева
	fmt.Println("--- БИНАРНОЕ ДЕРЕВО ---")
	tree := NewBinaryTree()
	
	// Добавляем элементы
	values := []int{50, 30, 70, 20, 40, 60, 80}
	fmt.Printf("Добавляем значения: %v\n", values)
	
	for _, v := range values {
		tree.Insert(v)
	}
	
	// Обходы дерева
	tree.InOrder()
	tree.PreOrder()
	tree.PostOrder()
	
	// Поиск элементов
	fmt.Println("\nПоиск элементов:")
	fmt.Printf("Поиск 40: %v\n", tree.Search(40))
	fmt.Printf("Поиск 90: %v\n", tree.Search(90))
	fmt.Printf("Поиск 50: %v\n", tree.Search(50))
}