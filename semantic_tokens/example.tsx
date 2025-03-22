import React, { useState, useEffect, FC } from 'react';
import axios from 'axios';

// 接口定义
interface User {
  id: number;
  name: string;
  email: string;
  isActive: boolean;
}

// 常量
const API_URL = 'https://api.example.com/users';
const MAX_USERS = 10;

// 枚举
enum UserRole {
  ADMIN = 'admin',
  EDITOR = 'editor',
  VIEWER = 'viewer',
}

/**
 * 用户列表组件
 * @param props - 组件属性
 */
const UserList: FC<{
  initialRole?: UserRole;
  onUserSelect?: (user: User) => void;
}> = ({ initialRole = UserRole.VIEWER, onUserSelect }) => {
  // 状态hooks
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null | undefined>(null);
  const [role, setRole] = useState<UserRole>(initialRole);

  // 副作用hook
  useEffect(() => {
    // 异步函数
    const fetchUsers = async (): Promise<void> => {
      try {
        setLoading(true);
        const response = await axios.get<User[]>(API_URL);
        // 数组方法和条件表达式
        const filteredUsers = response.data.filter((user) => user.isActive).slice(0, MAX_USERS);
        setUsers(filteredUsers);
        setError(null);
      } catch (err) {
        setError(`获取用户失败: ${err instanceof Error ? err.message : String(err)}`);
        setUsers([]);
      } finally {
        setLoading(false);
      }
    };

    fetchUsers();

    // 清理函数
    return () => {
      console.log('组件卸载');
    };
  }, []);

  // 事件处理函数
  const handleUserClick = (user: User): void => {
    console.log(`选择用户: ${user.name}`);
    if (onUserSelect) {
      onUserSelect(user);
    }
  };

  // JSX中的条件渲染和列表渲染
  return (
    <div className='user-list-container'>
      <h2>用户列表 ({role})</h2>

      {/* 角色选择器 */}
      <select value={role} onChange={(e) => setRole(e.target.value as UserRole)}>
        {Object.values(UserRole).map((roleValue) => (
          <option key={roleValue} value={roleValue}>
            {roleValue.charAt(0).toUpperCase() + roleValue.slice(1)}
          </option>
        ))}
      </select>

      {loading && <p>加载中...</p>}

      {error && <div className='error-message'>错误: {error}</div>}

      {!loading && !error && (
        <ul>
          {users.length > 0 ? (
            users.map((user) => (
              <li key={user.id} onClick={() => handleUserClick(user)} className={user.isActive ? 'active' : 'inactive'}>
                <span>{user.name}</span>
                <small>{user.email}</small>
              </li>
            ))
          ) : (
            <p>没有可用的用户</p>
          )}
        </ul>
      )}

      {/* 模板字符串和表达式插值 */}
      <p className='summary'>{`共有 ${users.length} 个用户，角色为 ${role}`}</p>
    </div>
  );
};

// 高阶组件示例
function withLogger<T extends object>(Component: React.ComponentType<T>): React.FC<T> {
  return (props: T) => {
    console.log('组件渲染', props);
    return <Component {...props} />;
  };
}

// 导出增强组件
export default withLogger(UserList);

// 工具类型
type Nullable<T> = T | null;
type UserPartial = Partial<User>;

// 工具函数
export const formatUser = (user: User): string => {
  return `${user.name} (${user.email})`;
};
