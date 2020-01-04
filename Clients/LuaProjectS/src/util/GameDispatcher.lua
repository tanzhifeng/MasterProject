local GameDispatcher = class("GameDispatcher")

function GameDispatcher:ctor()
	self.handlers_ = {}
end

function GameDispatcher:addEventListener(eventName, object, callback)
	self.handlers_[eventName] = self.handlers_[eventName] or {}
	local listener = self.handlers_[eventName]
	listener[object] = callback
end

function GameDispatcher:removeEventListener(eventName, object)
	if self.handlers_[eventName] and self.handlers_[eventName][object] then
		self.handlers_[eventName][object] = nil
	end
end

function GameDispatcher:dispatchEvent(eventName, param)
	if self.handlers_[eventName] == nil then
		return
	end
	local listener = clone(self.handlers_[eventName])
	for k, v in pairs(listener) do
		v(k, param)           
	end
end

return GameDispatcher


