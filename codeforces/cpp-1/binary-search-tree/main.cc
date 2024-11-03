#include<iostream>
#include<string>

template<typename Key, typename Value>
struct TreeNode {
    TreeNode* left;
    TreeNode* right;
    TreeNode* p;

    Key key;
    Value value;

    TreeNode(TreeNode* l, TreeNode* r, TreeNode* p): left(l), right(r), p(p) {}
};

class BSTree {
private:
    TreeNode<int, std::string>* head;
public:
    BSTree(): head(new TreeNode<int, std::string>(nullptr, nullptr, nullptr)) {}
};

int main() {
    BSTree tree;
    std::cout << "Hello World!" << std::endl;
}
