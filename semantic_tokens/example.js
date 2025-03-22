// Variables and constants
const PI = 3.14; // constant
let radius = 5; // variable

// Functions
function calculateArea(r) {
  // Parameters and return
  return PI * r * r; // function call, constant, parameter
}

// Classes and methods
class Circle {
  constructor(radius) {
    this.radius = radius; // property
  }

  getArea() {
    return calculateArea(this.radius); // method call, property
  }
}

// Objects and properties
const myCircle = new Circle(radius); // object instantiation
console.log('Area:', myCircle.getArea()); // console, method call

// Control flow
if (radius > 0) {
  // condition
  console.log('Valid radius'); // console
} else {
  console.error('Invalid radius'); // console
}

// Loops
for (let i = 0; i < 5; i++) {
  // loop
  console.log('Iteration:', i); // console, variable
}

// Arrays
const numbers = [1, 2, 3];
numbers.forEach((num) => console.log(num)); // array, function call

// Promises and async/await
async function fetchData(url) {
  try {
    const response = await fetch(url); // async/await, function call
    const data = await response.json(); // method call
    console.log(data); // console
  } catch (error) {
    console.error(error); // console
  }
}

fetchData();

// Regular expressions
const regex = /[A-Za-z]+/g; // regex
const text = 'Hello World';
const matches = text.match(regex); // method call
console.log(matches); // console

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
  email,
  address: { city },
} = user;

console.log(email, city);

// 数组与数组解构
const colors = ['red', 'green', 'blue'];
const [primaryColor, secondaryColor, ...otherColors] = colors;
console.log(primaryColor, secondaryColor, otherColors);

// 箭头函数
const double = (x) => x * 2;
const greet = (name) => {
  console.log(`Hello, ${name}!`);
  return `Greeting sent to ${name}`;
};
console.log(double(6), greet('foo'));

// 回调函数
setTimeout(() => {
  console.log('Timeout executed');
}, 1000);

// 类声明
class Product {
  constructor(name, price) {
    this.name = name;
    this.price = price;
    this._discount = 0;
  }

  get discount() {
    return this.price * (1 - this._discount);
  }

  set discount(value) {
    this._discount = value > 0 && value < 1 ? value : 0;
  }

  applyDiscount(percentage) {
    this._discount = percentage / 100;
    return this.discount;
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
const dp = new DigitalProduct();
console.log(dp);

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

// 正则表达式
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
const isValidEmail = emailRegex.test('user@example.com');
console.log(isValidEmail);

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
createCounter();

// Symbol
const uniqueKey = Symbol('description');
const obj = {
  [uniqueKey]: 'This property uses a Symbol as a key',
};
console.log(obj);

// Set
const uniqueItems = new Set([1, 2, 3, 1, 2]);
console.log(uniqueItems);

// Map
const userRoles = new Map();
userRoles.set(1001, ['admin']);
userRoles.set(1002, ['editor', 'viewer']);
