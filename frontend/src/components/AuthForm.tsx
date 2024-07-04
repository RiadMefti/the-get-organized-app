import React, { useState } from "react";
import { LockOutlined, MailOutlined } from "@ant-design/icons";
import { Button, Form, Input, Typography } from "antd";
import { useNavigate } from "react-router-dom";

interface AuthFormProps {
  onSubmit: (email: string, password: string, copyPassword?: string) => void;
  buttonText: string;
}

const AuthForm: React.FC<AuthFormProps> = ({ onSubmit, buttonText }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [copyPassword, setCopyPassword] = useState("");

  const { Title } = Typography;

  const navigate = useNavigate();

  const handleSubmit = (values: any) => {
    if (buttonText === "Register") {
      onSubmit(values.email, values.password, values.copyPassword);
    } else {
      onSubmit(values.email, values.password);
    }
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        padding:'8rem',
        alignItems: "center",
      }}
    >
      <Title level={2}>{buttonText}</Title>
      <Form
        name="auth_form"
        className="auth-form"
        initialValues={{ remember: true }}
        onFinish={handleSubmit}
        style={{
          width: "33%",
        }}
      >
        <Form.Item
          name="email"
          rules={[{ required: true, message: "Please input your Email!" }]}
        >
          <Input
            prefix={<MailOutlined className="site-form-item-icon" />}
            placeholder="Email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[{ required: true, message: "Please input your Password!" }]}
        >
          <Input
            prefix={<LockOutlined className="site-form-item-icon" />}
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </Form.Item>
        {buttonText === "Register" && (
          <Form.Item
            name="copyPassword"
            rules={[
              { required: true, message: "Please confirm your Password!" },
            ]}
          >
            <Input
              prefix={<LockOutlined className="site-form-item-icon" />}
              type="password"
              placeholder="Confirm Password"
              value={copyPassword}
              onChange={(e) => setCopyPassword(e.target.value)}
              required
            />
          </Form.Item>
        )}

        <Form.Item>
          <Button type="primary" htmlType="submit" className="auth-form-button">
            {buttonText}
          </Button>
          {buttonText !== "Register" ? (
            <span>
              {" "}
              Or <a onClick={() => navigate("/register")}>register now!</a>
            </span>
          ) : (
            <span>
              {" "}
              Or <a onClick={() => navigate("/login")}>Login now!</a>
            </span>
          )}
        </Form.Item>
      </Form>
    </div>
  );
};

export default AuthForm;
