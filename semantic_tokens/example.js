// 变量声明与基本数据类型
const username = 'alice'; // 字符串字面量
let age = 28; // 数字字面量
var isActive = true; // 布尔字面量
const PI = 3.14159; // 常量
let nullValue = null; // null值
let undefinedValue; // undefined值

// 对象与解构
const user = {
  id: 1001,
  name: 'Bob Smith',
  email: 'bob@example.com',
  address: {
    city: 'New York',
    zipCode: '10001',
  },
  roles: ['admin', 'editor'],
};

// 对象解构
const {
  name,
  email,
  address: { city },
} = user;

console.log(name, email, city);

// 数组与数组解构
const colors = ['red', 'green', 'blue'];
const [primaryColor, secondaryColor, ...otherColors] = colors;

// 函数声明
function calculateTotal(price, quantity, taxRate = 0.1) {
  return price * quantity * (1 + taxRate);
}

// 箭头函数
const double = (x) => x * 2;
const greet = (name) => {
  console.log(`Hello, ${name}!`);
  return `Greeting sent to ${name}`;
};

// 回调函数
setTimeout(() => {
  console.log('Timeout executed');
}, 1000);

// 高阶函数
const numbers = [1, 2, 3, 4, 5];
const doubled = numbers.map((num) => num * 2);
const evenNumbers = numbers.filter((num) => num % 2 === 0);
const sum = numbers.reduce((total, current) => total + current, 0);

// 类声明
class Product {
  constructor(name, price) {
    this.name = name;
    this.price = price;
    this._discount = 0;
  }

  get discountedPrice() {
    return this.price * (1 - this._discount);
  }

  set discount(value) {
    this._discount = value > 0 && value < 1 ? value : 0;
  }

  applyDiscount(percentage) {
    this._discount = percentage / 100;
    return this.discountedPrice;
  }

  static compare(productA, productB) {
    return productA.price - productB.price;
  }
}

// 继承
class DigitalProduct extends Product {
  constructor(name, price, downloadUrl) {
    super(name, price);
    this.downloadUrl = downloadUrl;
    this.type = 'digital';
  }

  deliver() {
    console.log(`Product ${this.name} can be downloaded at ${this.downloadUrl}`);
  }
}

// 异步编程 - Promise
function fetchUserData(userId) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (userId > 0) {
        resolve({ id: userId, name: 'Sample User' });
      } else {
        reject(new Error('Invalid user ID'));
      }
    }, 1000);
  });
}

// Promise 链
fetchUserData(123)
  .then((userData) => {
    console.log('User data:', userData);
    return userData.id;
  })
  .then((userId) => {
    return fetchUserData(userId + 1);
  })
  .catch((error) => {
    console.error('Error fetching user:', error.message);
  })
  .finally(() => {
    console.log('Operation completed');
  });

// Async/Await
async function loadUserProfile(userId) {
  try {
    const userData = await fetchUserData(userId);
    const posts = await fetchUserPosts(userData.id);
    return { user: userData, posts };
  } catch (error) {
    console.error('Failed to load profile:', error);
    throw error;
  }
}

// 模块导出
export { Product, DigitalProduct };
export default calculateTotal;

// DOM 操作
const button = document.getElementById('submit-button');
button.addEventListener('click', function (event) {
  event.preventDefault();
  const formData = new FormData(document.querySelector('form'));
  console.log('Form submitted:', Object.fromEntries(formData));
});

// 正则表达式
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
const isValidEmail = emailRegex.test('user@example.com');

// 模板字符串
const greeting = `Welcome, ${username}!
Your account is ${isActive ? 'active' : 'inactive'}.
Your current level: ${calculateLevel(age, isActive)}`;

// 闭包
function createCounter() {
  let count = 0;
  return {
    increment() {
      count++;
      return count;
    },
    decrement() {
      count--;
      return count;
    },
    getCount() {
      return count;
    },
  };
}

// Symbol
const uniqueKey = Symbol('description');
const obj = {
  [uniqueKey]: 'This property uses a Symbol as a key',
};

// Set 和 Map
const uniqueItems = new Set([1, 2, 3, 1, 2]);
const userRoles = new Map();
userRoles.set(1001, ['admin']);
userRoles.set(1002, ['editor', 'viewer']);
