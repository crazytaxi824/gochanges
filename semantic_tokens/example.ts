// 类型定义与接口
interface User {
  id: number;
  name: string;
  age?: number;
  isActive: boolean;
}
const user: User = { id: 101, name: 'John', age: 30, isActive: true }; // Example User object

// 枚举类型
enum Role {
  Admin = 'ADMIN',
  User = 'USER',
  Guest = 'GUEST',
}

type Foo = null | undefined;
let myFoo: Foo = null; // or `undefined`
console.log(myFoo);

// 泛型类
class DataStore<T> {
  private data: T[];

  constructor(initialData: T[] = []) {
    this.data = initialData;
  }

  add(item: T): void {
    this.data.push(item);
  }

  getAll(): ReadonlyArray<T> {
    return this.data;
  }
}

// 装饰器
function log(target: any, propertyKey: string, descriptor: PropertyDescriptor) {
  console.log(target);

  const originalMethod = descriptor.value;

  descriptor.value = function (...args: any[]) {
    console.log(`Calling ${propertyKey} with args: ${JSON.stringify(args)}`);
    return originalMethod.apply(this, args);
  };

  return descriptor;
}

log(null, 'omg', {});

// 使用装饰器的类
class UserService {
  private users: User[] = [];

  //@log
  addUser(user: User): void {
    this.users.push(user);
  }

  findUserById(id: number): User | undefined {
    return this.users.find((user) => user.id === id);
  }
}

// 异步函数与Promise
async function fetchUserData(userId: number): Promise<User> {
  try {
    const response = await fetch(`https://api.example.com/users/${userId}`);
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    return (await response.json()) as User;
  } catch (error) {
    console.error(`Failed to fetch user: ${error instanceof Error ? error.message : String(error)}`);
    throw error;
  }
}
fetchUserData(1);

// 类型断言与字面量类型
type Direction = 'north' | 'south' | 'east' | 'west';
let currentDirection: Direction = 'north';
console.log(currentDirection);

// 映射类型与条件类型
type Readonly<T> = {
  readonly [P in keyof T]: T[P];
};

type UserReadOnly = Readonly<User>;
const userReadOnly: UserReadOnly = Object.freeze(user); // Make it readonly
console.log(userReadOnly);

// 工具类型
type Nullable<T> = T | null;
type UserOrNull = Nullable<User>;
const nullUser: UserOrNull = null;
console.log(nullUser);

// 类型交叉与联合
type AdminUser = User & { permissions: string[] };
type Entity = User | { productId: string; price: number };
const adminUser: AdminUser = {
  id: 1,
  name: 'admin',
  age: 24,
  isActive: true,
  permissions: ['read', 'write'],
};
console.log(adminUser);

// 使用类型守卫的函数
function processEntity(entity: Entity): void {
  if ('name' in entity) {
    console.log(`Processing user: ${entity.name}`);
  } else {
    console.log(`Processing product: ${entity.productId}`);
  }
}
processEntity(userReadOnly);

// 与 DOM 交互
const button = document.getElementById('submit-btn') as HTMLButtonElement;
button?.addEventListener('click', (event: MouseEvent) => {
  event.preventDefault();
  console.log('Button clicked!');
});

// 命名空间
namespace Validation {
  export interface StringValidator {
    isValid(s: string): boolean;
  }

  export class EmailValidator implements StringValidator {
    isValid(email: string): boolean {
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      return emailRegex.test(email);
    }
  }
}
let validator = new Validation.EmailValidator();
console.log(validator.isValid('test@example.com')); // true or false

// 模块导入导出
export { User, Role, UserService };
export default DataStore;

class Animal {
  constructor(public name: string) {}
}

class Dog extends Animal {
  constructor(
    name: string,
    public breed: string,
  ) {
    super(name); // 调用父类的构造函数
    // 此处才能使用 this
  }
}
const dog = new Dog('doggy', 'breed');
console.log(dog);
