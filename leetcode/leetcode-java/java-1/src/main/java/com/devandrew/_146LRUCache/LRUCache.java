package com.devandrew._146LRUCache;

import java.util.HashMap;
import java.util.Map;

public class LRUCache {
    class Node {
        private int key;
        private int val;
        private Node next;
        private Node prev;

        public Node() {

        }

        public Node(int key, int val) {
            this.key = key;
            this.val = val;
        }
    }

    private Map<Integer, Node> nodes;
    private Node head;
    private Node tail;
    private int capacity;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.nodes = new HashMap<>();
        this.head = new Node();
        this.tail = new Node();

        head.prev = tail;
        tail.next = head;
    }

    public int get(int key) {
        Node node = nodes.get(key);

        if (node != null) {
            moveToHead(node);
            return node.val;
        }

        return -1;
    }

    public void put(int key, int value) {
        Node node = nodes.get(key);

        if (node != null) {
            node.val = value;
            moveToHead(node);
        } else {
            if (nodes.size() == capacity) {
                removeTail();
            }

            Node newNode = new Node(key, value);
            insertToHead(newNode);
            nodes.put(key, newNode);
        }
    }

    private void moveToHead(Node node) {
        if (node != head.prev) {
            removeNode(node);
            insertToHead(node);
        }
    }

    private void removeNode(Node node) {
        Node prev = node.prev;
        Node next = node.next;

        prev.next = next;
        next.prev = prev;
    }

    private void removeTail() {
        nodes.remove(tail.next.key);
        removeNode(tail.next);
    }

    private void insertToHead(Node node) {
        Node prev = head.prev;

        head.prev = node;
        prev.next = node;

        node.next = head;
        node.prev = prev;
    }
}
