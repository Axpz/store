"use client";

import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import { toast } from "react-toastify";
import { useAuth } from "@/context/UserContext";
import { Separator } from "@/components/ui/separator";

export default function UserManagementSettings() {
  const { user } = useAuth();
  const [newUsername, setNewUsername] = useState("");
  const [oldPassword, setOldPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmNewPassword, setConfirmNewPassword] = useState("");

  const handleUpdateUsername = () => {
    console.log("Updating username to:", newUsername);
    toast.success(`Username updated to: ${newUsername}`);
    setNewUsername("");
  };

  const handleUpdatePassword = () => {
    if (newPassword !== confirmNewPassword) {
      toast.error("New passwords do not match.");
      return;
    }
    console.log("Updating password:", { oldPassword, newPassword });
    toast.success("Password updated successfully.");
    setOldPassword("");
    setNewPassword("");
    setConfirmNewPassword("");
  };

  return (
    <div className="flex-1 bg-white dark:bg-gray-800 shadow-lg rounded-lg p-8">
      <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">
        User Management
      </h2>

      <div className="space-y-6">
        {/* Username Update */}
        <Card>
          <CardHeader>
            <CardTitle className="text-md font-semibold">Username</CardTitle>
          </CardHeader>
          <CardContent className="space-y-3">
            <div className="space-y-1">
              <Label htmlFor="new-username" className="text-sm text-gray-500">New Username</Label>
              <Input
                id="new-username"
                type="text"
                value={newUsername}
                onChange={(e) => setNewUsername(e.target.value)}
                className="w-full rounded-md"
              />
            </div>
            <Button onClick={handleUpdateUsername} className="w-full bg-blue-50 text-blue-500 hover:bg-blue-100 font-semibold py-2 rounded-md">
              Update Username
            </Button>
          </CardContent>
        </Card>

        {/* Password Update */}
        <Card>
          <CardHeader>
            <CardTitle className="text-md font-semibold">Password</CardTitle>
          </CardHeader>
          <CardContent className="space-y-3">
            <div className="space-y-1">
              <Label htmlFor="old-password" className="text-sm text-gray-500">Old Password</Label>
              <Input
                id="old-password"
                type="password"
                value={oldPassword}
                onChange={(e) => setOldPassword(e.target.value)}
                className="w-full rounded-md"
              />
            </div>
            <div className="space-y-1">
              <Label htmlFor="new-password" className="text-sm text-gray-500">New Password</Label>
              <Input
                id="new-password"
                type="password"
                value={newPassword}
                onChange={(e) => setNewPassword(e.target.value)}
                className="w-full rounded-md"
              />
            </div>
            <div className="space-y-1">
              <Label htmlFor="confirm-new-password" className="text-sm text-gray-500">Confirm New Password</Label>
              <Input
                id="confirm-new-password"
                type="password"
                value={confirmNewPassword}
                onChange={(e) => setConfirmNewPassword(e.target.value)}
                className="w-full rounded-md"
              />
            </div>
            <Button
              onClick={handleUpdatePassword}
              className="w-full bg-blue-50 text-blue-500 hover:bg-blue-100 font-semibold py-2 rounded-md"
            >
              Update Password
            </Button>
          </CardContent>
        </Card>

        {/* User Information */}
        <Card>
          <CardHeader>
            <CardTitle className="text-md font-semibold">Information</CardTitle>
          </CardHeader>
          <CardContent className="space-y-3">
            <div className="space-y-1">
              <Label className="text-sm text-gray-500">Current Username</Label>
              <p className="text-sm font-medium">{user?.username || 'Not logged in'}</p>
            </div>
            <div className="space-y-1">
              <Label className="text-sm text-gray-500">Email</Label>
              <p className="text-sm font-medium">{user?.email || 'Not available'}</p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}