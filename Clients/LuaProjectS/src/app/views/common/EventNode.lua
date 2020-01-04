EventNode = class("EventNode", function()
		local node = display.newNode()
		node:enableNodeEvents()
		return node
	end)

function EventNode:ctor()
end

function EventNode:close()
	self:removeSelf()
end